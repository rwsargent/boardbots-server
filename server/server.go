package server

import (
	pb "boardbots-server/bbpb"
	"boardbots-server/manager"
	"boardbots-server/server/persistence"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc/metadata"
)

type ServerMode int

const (
	Development ServerMode = iota
	Production
)

type Server struct {
	gameManager manager.GameManager
	userPortal  persistence.UserPortal
	mode        ServerMode
}

func (server *Server) GetGames(ctx context.Context, request *pb.GameRequest) (*pb.GameResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		tokens := md.Get("token")
		fmt.Printf("tokens: %v\n", tokens)
	}
	uuid := request.GameId
	reversed := reverse(uuid.GetValue())
	return &pb.GameResponse{
		GameId: &pb.UUID{Value: reversed},
	}, nil
}

func (server *Server) Authenticate(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	res := pb.AuthResponse{}
	if server.userPortal.ValidateCredentials(req.Username, req.Password) {
		res.Token = generateFakeToken()
	}
	return &res, nil
}

func generateFakeToken() string {
	randBytes := make([]byte, 128)
	_, err := rand.Read(randBytes)
	if err != nil {
		log.Errorf("could not generate rand token: %v\n", err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(randBytes)
}

func NewServer(mode ServerMode) *Server {
	userPortal := persistence.NewMemoryPortal()
	userPortal.NewUser("rws", "rws")
	return &Server{
		manager.NewMemoryGameManager(),
		userPortal,
		mode,
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

package server

import (
	pb "boardbots-server/bbpb"
	"context"
)

type Server struct{}

func (server *Server) GetGames(ctx context.Context, request *pb.GameRequest) (*pb.GameResponse, error) {
	uuid := request.GameId
	reversed := reverse(uuid.GetValue())
	return &pb.GameResponse{
		GameId: &pb.UUID{Value: reversed},
	}, nil
}

func NewServer() *Server {
 return &Server{}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: boardbots.proto

/*
Package boardbots is a generated protocol buffer package.

It is generated from these files:
	boardbots.proto

It has these top-level messages:
	UUID
	GameRequest
	PlayerState
	Position
	Piece
	GameResponse
	AuthRequest
	AuthResponse
*/
package boardbots

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Piece_Type int32

const (
	Piece_BARRIER Piece_Type = 0
	Piece_PAWN    Piece_Type = 1
)

var Piece_Type_name = map[int32]string{
	0: "BARRIER",
	1: "PAWN",
}
var Piece_Type_value = map[string]int32{
	"BARRIER": 0,
	"PAWN":    1,
}

func (x Piece_Type) String() string {
	return proto.EnumName(Piece_Type_name, int32(x))
}
func (Piece_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type UUID struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *UUID) Reset()                    { *m = UUID{} }
func (m *UUID) String() string            { return proto.CompactTextString(m) }
func (*UUID) ProtoMessage()               {}
func (*UUID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UUID) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GameRequest struct {
	GameId *UUID `protobuf:"bytes,1,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
}

func (m *GameRequest) Reset()                    { *m = GameRequest{} }
func (m *GameRequest) String() string            { return proto.CompactTextString(m) }
func (*GameRequest) ProtoMessage()               {}
func (*GameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GameRequest) GetGameId() *UUID {
	if m != nil {
		return m.GameId
	}
	return nil
}

type PlayerState struct {
	PlayerName   string    `protobuf:"bytes,1,opt,name=player_name,json=playerName" json:"player_name,omitempty"`
	PawnPosition *Position `protobuf:"bytes,2,opt,name=pawn_position,json=pawnPosition" json:"pawn_position,omitempty"`
	Barriers     int32     `protobuf:"varint,3,opt,name=barriers" json:"barriers,omitempty"`
}

func (m *PlayerState) Reset()                    { *m = PlayerState{} }
func (m *PlayerState) String() string            { return proto.CompactTextString(m) }
func (*PlayerState) ProtoMessage()               {}
func (*PlayerState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PlayerState) GetPlayerName() string {
	if m != nil {
		return m.PlayerName
	}
	return ""
}

func (m *PlayerState) GetPawnPosition() *Position {
	if m != nil {
		return m.PawnPosition
	}
	return nil
}

func (m *PlayerState) GetBarriers() int32 {
	if m != nil {
		return m.Barriers
	}
	return 0
}

type Position struct {
	Row int32 `protobuf:"varint,1,opt,name=row" json:"row,omitempty"`
	Col int32 `protobuf:"varint,2,opt,name=col" json:"col,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Position) GetRow() int32 {
	if m != nil {
		return m.Row
	}
	return 0
}

func (m *Position) GetCol() int32 {
	if m != nil {
		return m.Col
	}
	return 0
}

type Piece struct {
	Type     Piece_Type `protobuf:"varint,1,opt,name=type,enum=Piece_Type" json:"type,omitempty"`
	Position *Position  `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
	Owner    int32      `protobuf:"varint,3,opt,name=owner" json:"owner,omitempty"`
}

func (m *Piece) Reset()                    { *m = Piece{} }
func (m *Piece) String() string            { return proto.CompactTextString(m) }
func (*Piece) ProtoMessage()               {}
func (*Piece) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Piece) GetType() Piece_Type {
	if m != nil {
		return m.Type
	}
	return Piece_BARRIER
}

func (m *Piece) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Piece) GetOwner() int32 {
	if m != nil {
		return m.Owner
	}
	return 0
}

type GameResponse struct {
	GameId      *UUID                      `protobuf:"bytes,1,opt,name=game_id,json=gameId" json:"game_id,omitempty"`
	Players     []*PlayerState             `protobuf:"bytes,2,rep,name=players" json:"players,omitempty"`
	CurrentTurn int32                      `protobuf:"varint,3,opt,name=current_turn,json=currentTurn" json:"current_turn,omitempty"`
	StartDate   *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=start_date,json=startDate" json:"start_date,omitempty"`
	EndDate     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=end_date,json=endDate" json:"end_date,omitempty"`
	Winner      int32                      `protobuf:"varint,6,opt,name=winner" json:"winner,omitempty"`
	Board       []*Piece                   `protobuf:"bytes,7,rep,name=board" json:"board,omitempty"`
}

func (m *GameResponse) Reset()                    { *m = GameResponse{} }
func (m *GameResponse) String() string            { return proto.CompactTextString(m) }
func (*GameResponse) ProtoMessage()               {}
func (*GameResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GameResponse) GetGameId() *UUID {
	if m != nil {
		return m.GameId
	}
	return nil
}

func (m *GameResponse) GetPlayers() []*PlayerState {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *GameResponse) GetCurrentTurn() int32 {
	if m != nil {
		return m.CurrentTurn
	}
	return 0
}

func (m *GameResponse) GetStartDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *GameResponse) GetEndDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func (m *GameResponse) GetWinner() int32 {
	if m != nil {
		return m.Winner
	}
	return 0
}

func (m *GameResponse) GetBoard() []*Piece {
	if m != nil {
		return m.Board
	}
	return nil
}

type AuthRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *AuthRequest) Reset()                    { *m = AuthRequest{} }
func (m *AuthRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()               {}
func (*AuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AuthRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *AuthResponse) Reset()                    { *m = AuthResponse{} }
func (m *AuthResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()               {}
func (*AuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*UUID)(nil), "UUID")
	proto.RegisterType((*GameRequest)(nil), "GameRequest")
	proto.RegisterType((*PlayerState)(nil), "PlayerState")
	proto.RegisterType((*Position)(nil), "Position")
	proto.RegisterType((*Piece)(nil), "Piece")
	proto.RegisterType((*GameResponse)(nil), "GameResponse")
	proto.RegisterType((*AuthRequest)(nil), "AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "AuthResponse")
	proto.RegisterEnum("Piece_Type", Piece_Type_name, Piece_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BoardbotsService service

type BoardbotsServiceClient interface {
	GetGames(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameResponse, error)
	Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type boardbotsServiceClient struct {
	cc *grpc.ClientConn
}

func NewBoardbotsServiceClient(cc *grpc.ClientConn) BoardbotsServiceClient {
	return &boardbotsServiceClient{cc}
}

func (c *boardbotsServiceClient) GetGames(ctx context.Context, in *GameRequest, opts ...grpc.CallOption) (*GameResponse, error) {
	out := new(GameResponse)
	err := grpc.Invoke(ctx, "/BoardbotsService/GetGames", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boardbotsServiceClient) Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := grpc.Invoke(ctx, "/BoardbotsService/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BoardbotsService service

type BoardbotsServiceServer interface {
	GetGames(context.Context, *GameRequest) (*GameResponse, error)
	Authenticate(context.Context, *AuthRequest) (*AuthResponse, error)
}

func RegisterBoardbotsServiceServer(s *grpc.Server, srv BoardbotsServiceServer) {
	s.RegisterService(&_BoardbotsService_serviceDesc, srv)
}

func _BoardbotsService_GetGames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardbotsServiceServer).GetGames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BoardbotsService/GetGames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardbotsServiceServer).GetGames(ctx, req.(*GameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoardbotsService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoardbotsServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BoardbotsService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoardbotsServiceServer).Authenticate(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BoardbotsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BoardbotsService",
	HandlerType: (*BoardbotsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGames",
			Handler:    _BoardbotsService_GetGames_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _BoardbotsService_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "boardbots.proto",
}

func init() { proto.RegisterFile("boardbots.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0x5d, 0xda, 0x38, 0x71, 0xae, 0xdd, 0x2d, 0x88, 0x32, 0x8c, 0xe9, 0xd6, 0xce, 0xec, 0xa3,
	0x30, 0xa6, 0x42, 0xc6, 0x1e, 0xf6, 0x98, 0xd2, 0x52, 0xf2, 0x52, 0x82, 0x9a, 0xb2, 0xc7, 0xa0,
	0xc4, 0x77, 0xa9, 0x59, 0x2c, 0x79, 0x92, 0xdc, 0x90, 0x3d, 0xef, 0x8f, 0xec, 0x9f, 0x0e, 0x49,
	0x76, 0xc8, 0x1e, 0x46, 0xdf, 0x7c, 0xee, 0xe7, 0xf1, 0xb9, 0x47, 0xf0, 0x62, 0x21, 0xb9, 0xca,
	0x17, 0xd2, 0x68, 0x5a, 0x29, 0x69, 0x64, 0x7a, 0xba, 0x92, 0x72, 0xb5, 0xc6, 0x0b, 0x87, 0x16,
	0xf5, 0xf7, 0x0b, 0x53, 0x94, 0xa8, 0x0d, 0x2f, 0x2b, 0x5f, 0x90, 0x9d, 0x40, 0xf7, 0xfe, 0x7e,
	0x72, 0x45, 0x8e, 0x21, 0x78, 0xe4, 0xeb, 0x1a, 0x93, 0xce, 0x59, 0xe7, 0x7c, 0xc0, 0x3c, 0xc8,
	0x3e, 0x41, 0x74, 0xc3, 0x4b, 0x64, 0xf8, 0xb3, 0x46, 0x6d, 0xc8, 0x6b, 0xe8, 0xaf, 0x78, 0x89,
	0xf3, 0x22, 0x77, 0x65, 0xd1, 0x28, 0xa0, 0xb6, 0x99, 0xf5, 0x6c, 0x74, 0x92, 0x67, 0xbf, 0x20,
	0x9a, 0xae, 0xf9, 0x16, 0xd5, 0x9d, 0xe1, 0x06, 0xc9, 0x29, 0x44, 0x95, 0x83, 0x73, 0xc1, 0xcb,
	0x76, 0x32, 0xf8, 0xd0, 0x2d, 0x2f, 0x91, 0x50, 0x38, 0xaa, 0xf8, 0x46, 0xcc, 0x2b, 0xa9, 0x0b,
	0x53, 0x48, 0x91, 0x1c, 0xb8, 0xa9, 0x03, 0x3a, 0x6d, 0x02, 0x2c, 0xb6, 0xf9, 0x16, 0x91, 0x14,
	0xc2, 0x05, 0x57, 0xaa, 0x40, 0xa5, 0x93, 0xc3, 0xb3, 0xce, 0x79, 0xc0, 0x76, 0x38, 0xa3, 0x10,
	0xee, 0xea, 0x86, 0x70, 0xa8, 0xe4, 0xc6, 0x2d, 0x0c, 0x98, 0xfd, 0xb4, 0x91, 0xa5, 0x5c, 0xbb,
	0xf9, 0x01, 0xb3, 0x9f, 0xd9, 0xef, 0x0e, 0x04, 0xd3, 0x02, 0x97, 0x96, 0x66, 0xd7, 0x6c, 0x2b,
	0xcf, 0xef, 0xf9, 0x28, 0xa2, 0x2e, 0x4a, 0x67, 0xdb, 0x0a, 0x99, 0x4b, 0x90, 0x77, 0x10, 0xfe,
	0x9f, 0xe1, 0x2e, 0x65, 0x25, 0x94, 0x1b, 0x81, 0xaa, 0xa1, 0xe6, 0x41, 0xf6, 0x0a, 0xba, 0x76,
	0x14, 0x89, 0xa0, 0x7f, 0x39, 0x66, 0x6c, 0x72, 0xcd, 0x86, 0xcf, 0x48, 0x08, 0xdd, 0xe9, 0xf8,
	0xdb, 0xed, 0xb0, 0x93, 0xfd, 0x39, 0x80, 0xd8, 0x4b, 0xac, 0x2b, 0x29, 0x34, 0x3e, 0xa5, 0x31,
	0x79, 0x0f, 0x7d, 0xaf, 0xa0, 0x4e, 0x0e, 0xce, 0x0e, 0xcf, 0xa3, 0x51, 0x4c, 0xf7, 0x34, 0x67,
	0x6d, 0x92, 0xbc, 0x81, 0x78, 0x59, 0x2b, 0x85, 0xc2, 0xcc, 0x4d, 0xad, 0x44, 0x43, 0x2a, 0x6a,
	0x62, 0xb3, 0x5a, 0x09, 0xf2, 0x15, 0x40, 0x1b, 0xae, 0xcc, 0x3c, 0xe7, 0x06, 0x93, 0xae, 0xdb,
	0x96, 0x52, 0xef, 0x18, 0xda, 0x3a, 0x86, 0xce, 0x5a, 0xc7, 0xb0, 0x81, 0xab, 0xbe, 0xb2, 0xa7,
	0xfd, 0x02, 0x21, 0x8a, 0xdc, 0x37, 0x06, 0x4f, 0x36, 0xf6, 0x51, 0xe4, 0xae, 0xed, 0x25, 0xf4,
	0x36, 0x85, 0xb0, 0x1a, 0xf5, 0x1c, 0x9d, 0x06, 0x91, 0x13, 0x08, 0x9c, 0x73, 0x93, 0xbe, 0xfb,
	0xa5, 0x9e, 0xbf, 0x01, 0xf3, 0xc1, 0xec, 0x1a, 0xa2, 0x71, 0x6d, 0x1e, 0x5a, 0x17, 0xa6, 0x10,
	0xd6, 0x1a, 0xd5, 0x9e, 0xa7, 0x76, 0xd8, 0xe6, 0x2a, 0xae, 0xf5, 0x46, 0xaa, 0xdc, 0x9d, 0x6a,
	0xc0, 0x76, 0x38, 0x7b, 0x0b, 0xb1, 0x1f, 0xd3, 0x28, 0x7d, 0x0c, 0x81, 0x91, 0x3f, 0x50, 0xb4,
	0x96, 0x77, 0x60, 0xf4, 0x00, 0xc3, 0xcb, 0xf6, 0x11, 0xdd, 0xa1, 0x7a, 0x2c, 0x96, 0x48, 0x3e,
	0x40, 0x78, 0x83, 0xc6, 0x9e, 0x49, 0x93, 0x98, 0xee, 0xbd, 0x88, 0xf4, 0x88, 0xfe, 0x73, 0xbc,
	0x8f, 0x7e, 0x05, 0x0a, 0x53, 0x2c, 0xed, 0xff, 0xc6, 0x74, 0x8f, 0x78, 0x7a, 0x44, 0xf7, 0xf7,
	0x2f, 0x7a, 0x4e, 0xa9, 0xcf, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x32, 0x94, 0x11, 0xb5,
	0x03, 0x00, 0x00,
}

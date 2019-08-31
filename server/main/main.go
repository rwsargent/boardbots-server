package main

import (
	pb "boardbots-server/bbpb"
	"boardbots-server/server"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":8765")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBoardbotsServiceServer(s, server.NewServer(server.Development))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	shutdownGracefully(s)
}

func shutdownGracefully(grpc *grpc.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Printf("Shutting down")
	defer cancel()
	grpc.GracefulStop()
}

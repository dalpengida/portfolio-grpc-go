package main

import (
	"net"

	"github.com/dalpengida/portfolio-grpc-go/config"
	"github.com/dalpengida/portfolio-grpc-go/model/chat"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", config.Host)
	if err != nil {
		log.Error().Msgf("failed to listen: %v", err)
	}

	chatRpc := chat.New()
	s := grpc.NewServer()
	chat.RegisterChatRoomServer(s, &chatRpc)

	if err := s.Serve(lis); err != nil {
		log.Error().Msgf("failed to serve: %v", err)
	}
}

package main

import (
	"net"

	"github.com/dalpengida/portfolio-grpc-go/config"
	"github.com/dalpengida/portfolio-grpc-go/model/chat"
	"github.com/dalpengida/portfolio-grpc-go/model/health"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const ()

var ()

func init() {
	listen, err := net.Listen("tcp", config.Host)
	if err != nil {
		panic(err) // 못 열면 망가진 거임
	}
	defer listen.Close()

	healthRpc := health.New()
	chatRpc := chat.New()
	grpcServer := grpc.NewServer()
	health.RegisterHealthServer(
		grpcServer,
		&healthRpc,
	)

	chat.RegisterChatRoomServer(
		grpcServer,
		&chatRpc,
	)

	serviceInfo := grpcServer.GetServiceInfo()
	log.Info().Interface("info", serviceInfo).Msg("")

	grpcServer.Serve(listen)

}

func main() {

}

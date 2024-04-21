package main

import (
	"context"
	"os"
	"time"

	"github.com/dalpengida/portfolio-grpc-go/config"
	"github.com/dalpengida/portfolio-grpc-go/model/chat"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(config.Host, grpc.WithInsecure())
	if err != nil {
		log.Error().Err(err).Msg("dial failed")
		panic(err)
	}
	defer conn.Close()
	c := chat.NewChatRoomClient(conn)

	clientName := os.Args[1]

	ctx := context.TODO()
	stream, err := c.Chat(ctx)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	go func() {
		for {
			mes := "hello, I'm " + clientName
			if err := stream.SendMsg(&chat.ChatRequest{Message: mes}); err != nil {
				log.Error().Err(err).Msg("")
			}
			log.Printf("sent: %s", mes)
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Error().Err(err).Msg("")
		}

		log.Info().Interface("resp", resp).Msg("")

	}
}

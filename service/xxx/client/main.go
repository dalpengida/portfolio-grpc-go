package main

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/dalpengida/portfolio-grpc-go/config"
	"github.com/dalpengida/portfolio-grpc-go/model/health"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ()

func init() {}

func main() {
	conn, err := grpc.Dial(
		config.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Err(err).Msg("dial failed")
		panic(err) // 연결 못하면 망함
	}
	defer conn.Close()

	client := health.NewHealthClient(conn)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		r, err := client.Health(context.TODO(), &health.HealthRequest{
			Id: uuid.NewString(),
		})
		if err != nil {
			log.Error().Err(err).Msg("health failed")
			panic(err)
		}

		for {
			time.Sleep(1 * time.Second)
			recv, err := r.Recv()
			if err != nil {
				if err == io.EOF {
					log.Info().Msg("EOF")
				} else {
					log.Error().Err(err).Msg("health failed")
					panic(err)
				}
			}

			log.Info().Interface("recv", recv).Msg("health success")
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)

			r, err := client.SayHello(context.TODO(), &health.HelloRequest{
				Name: uuid.NewString(),
			})
			if err != nil {
				log.Error().Err(err).Msg("say hello failed")
			}

			log.Info().Interface("result", r).Msg("sayho")

		}
	}()

	wg.Wait() // go 루틴으로 좀 테스트를 하기 위하여
}

package health

import (
	context "context"

	"github.com/rs/zerolog/log"
)

type healthRpc struct {
	UnimplementedHealthServer
	client           map[string]Health_HealthServer
	cnt              int
	broadcastChannel chan string
}

func New() healthRpc {
	return healthRpc{
		client:           make(map[string]Health_HealthServer),
		broadcastChannel: make(chan string),
	}
}

func (s *healthRpc) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	s.cnt++
	if s.cnt%2 == 0 {
		s.broadcastChannel <- "broadcast"
	}

	log.Info().Interface("req", req).Interface("counter", s.cnt).Msg("sayhello")
	return &HelloResponse{
		Name: req.Name,
	}, nil
}

func (s *healthRpc) Health(req *HealthRequest, stream Health_HealthServer) error {
	if _, ok := s.client[req.Id]; !ok {
		s.client[req.Id] = stream
	}

	for msg := range s.broadcastChannel {
		for id, v := range s.client {
			err := v.Send(&HealthResponse{
				Id:  id,
				Val: msg,
			})
			if err != nil {
				log.Error().Err(err).Msg("send failed ")
			}
		}
	}

	return nil
}

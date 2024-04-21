package chat

import (
	"os"
	sync "sync"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type chatServer struct {
	UnimplementedChatRoomServer

	clients map[string]ChatRoom_ChatServer
	mu      sync.RWMutex
}

func New() chatServer {
	return chatServer{
		clients: make(map[string]ChatRoom_ChatServer),
		mu:      sync.RWMutex{},
	}
}

func (s *chatServer) addClient(uid string, srv ChatRoom_ChatServer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[uid] = srv
}

func (s *chatServer) removeClient(uid string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, uid)
}

func (s *chatServer) getClients() []ChatRoom_ChatServer {
	var cs []ChatRoom_ChatServer

	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, c := range s.clients {
		cs = append(cs, c)
	}
	return cs
}

func (s *chatServer) Chat(srv ChatRoom_ChatServer) error {
	uid := uuid.Must(uuid.NewRandom()).String()
	log.Printf("new user: %s", uid)

	s.addClient(uid, srv)
	defer s.removeClient(uid)

	defer func() {
		if err := recover(); err != nil {
			log.Error().Msgf("panic: %v", err)
			os.Exit(1)
		}
	}()

	for {
		resp, err := srv.Recv()
		if err != nil {
			log.Error().Err(err).Msg("")
			break
		}

		log.Info().Interface("resp", resp.Message).Msg("recv message")
		for _, ss := range s.getClients() {
			if err := ss.Send(&ChatResponse{Message: resp.Message}); err != nil {
				log.Error().Err(err).Msg("")
			}
		}
	}
	return nil
}

package metadata

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) mustEmbedUnimplementedMetadataServer() {
	panic("implement me")
}

func (s *Server) GetMetadata(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.CorrelationId)
	return &Message{CorrelationId: "TEST01"}, nil
}
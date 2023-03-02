package grpc

import (
	"go.mongodb.org/mongo-driver/mongo"
	"ms-workspace/ms-notification/global/config"
	"ms-workspace/ms-notification/internal/v1/repository"
	"ms-workspace/ms-notification/internal/v1/usecase"
	"ms-workspace/package/proto/ms-notification/v1/go-proto"
)

type Server struct {
	go_proto.UnimplementedEmailServiceServer
	mongoDb *mongo.Database

	emailRepo    repository.EmailRepository
	emailUseCase usecase.EmailUseCase
}

func NewServer(database *mongo.Database) *Server {
	server := &Server{}

	// repository
	server.emailRepo = repository.NewEmailRepository(database)

	// service
	server.emailUseCase = usecase.NewEmailUseCase(server.emailRepo, config.SendGridAPIKey, config.SendgridSenderName, config.SendgridSenderEmail)
	return server
}

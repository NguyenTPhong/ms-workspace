package grpc

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"ms-workspace/ms-user/internal/v1/repository"
	"ms-workspace/ms-user/internal/v1/usecase"
	go_proto "ms-workspace/package/proto/ms-user/v1/go-proto"
)

type Server struct {
	go_proto.UnimplementedUserServiceServer

	db    *gorm.DB
	redis *redis.Client

	authRepo      repository.AuthRepository
	authCacheRepo repository.AuthCacheRepository
	authUseCase   usecase.AuthUseCase
}

func NewServer(db *gorm.DB, redis *redis.Client) *Server {
	server := &Server{
		db:    db,
		redis: redis,
	}

	// init repository
	server.authRepo = repository.NewAuthRepository(server.db)
	server.authCacheRepo = repository.NewAuthCacheRepository(server.redis)

	// init use case
	server.authUseCase = usecase.NewUserUseCase(server.authRepo, server.authCacheRepo)

	return server
}

package main

import (
	"go.uber.org/zap"
	"log"
	"ms-workspace/ms-user/global/config"
	"ms-workspace/ms-user/internal/v1/repository/migration"
	"ms-workspace/package/db"
	"ms-workspace/package/redis"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"ms-workspace/ms-user/global"
	_grpc "ms-workspace/ms-user/internal/v1/delivery/grpc"
	pb "ms-workspace/package/proto/ms-user/v1/go-proto"
)

func main() {
	global.Init()
	defer global.DeInit()

	database, err := db.NewDatabase(config.DbConnStr, int(config.DbMaxConn), int(config.DbMaxIdleConn), int(config.DBLogLevel))
	if err != nil {
		panic(err)
	}
	defer db.Close(database)

	migration.CreateTable(database)

	redisClient, err := redis.NewClient(config.RedisHost, config.RedisPassword)
	if err != nil {
		panic(err)
	}
	defer redis.Close(redisClient)

	listenPort, err := net.Listen("tcp", ":81")
	if err != nil {
		global.Logger.Error("failed to listen", zap.Error(err))
		panic("server error")
	}

	s := grpc.NewServer()
	handlerServer := _grpc.NewServer(database, redisClient)
	pb.RegisterUserServiceServer(s, handlerServer)
	reflection.Register(s)

	if err = s.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

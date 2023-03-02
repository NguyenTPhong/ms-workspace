package main

import (
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"ms-workspace/ms-notification/global"
	"ms-workspace/ms-notification/global/config"
	_grpc "ms-workspace/ms-notification/internal/v1/delivery/grpc"
	"ms-workspace/package/db"
	pb "ms-workspace/package/proto/ms-notification/v1/go-proto"
)

func main() {
	global.Init()
	defer global.DeInit()

	mongoDb, cancelFunc, err := db.NewMongoDB(config.MongoConnStr, config.DBName)
	if err != nil {
		panic(err)
	}
	defer cancelFunc()

	listenPort, err := net.Listen("tcp", ":82")
	if err != nil {
		global.Logger.Error("failed to listen", zap.Error(err))
		panic("server error")
	}

	s := grpc.NewServer()
	handlerServer := _grpc.NewServer(mongoDb)
	pb.RegisterEmailServiceServer(s, handlerServer)
	reflection.Register(s)

	if err = s.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

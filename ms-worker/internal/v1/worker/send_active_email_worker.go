package worker

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"ms-workspace/ms-worker/global"
	"ms-workspace/ms-worker/global/config"
	"ms-workspace/package/grpc"
	go_proto "ms-workspace/package/proto/ms-notification/v1/go-proto"
)

type SendActiveEmailWorker struct {
	redis *redis.Client
}

func NewSendActiveEmailWorker(redis *redis.Client) *SendActiveEmailWorker {
	return &SendActiveEmailWorker{redis: redis}
}

func (w *SendActiveEmailWorker) Run() {
	// subscribe to the channel
	pubSub := w.redis.Subscribe(config.SendActiveEmailTopic)

	// wait for subscription confirmation
	_, err := pubSub.Receive()
	if err != nil {
		panic(err)
	}

	// start receiving messages
	ch := pubSub.Channel()
	for msg := range ch {
		go func(payload string) {
			err := w.HandleMessage(payload)
			if err != nil {
				global.Logger.Error("handle send active email message error", zap.Error(err))
			}
		}(msg.Payload)
	}
}

func (w *SendActiveEmailWorker) HandleMessage(payload string) error {
	var req go_proto.SendActiveEmailRequest
	err := json.Unmarshal([]byte(payload), &req)
	if err != nil {
		return err
	}

	conn, cancelFunc, err := grpc.ConnectToGrpcServer(config.NotificationServiceHost)
	if err != nil {
		global.Logger.Error("connect to notification service error", zap.Error(err))
	}

	defer cancelFunc()

	client := go_proto.NewEmailServiceClient(conn)

	res, err := client.SendActiveEmail(context.Background(), &req)

	if err != nil {
		return err
	}

	global.Logger.Info("send active email success", zap.Any("res", res))
	return nil
}

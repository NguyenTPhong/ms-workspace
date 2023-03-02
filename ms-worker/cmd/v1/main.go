package main

import (
	"ms-workspace/ms-worker/global"
	"ms-workspace/ms-worker/global/config"
	"ms-workspace/ms-worker/internal/v1/worker"
	"ms-workspace/package/redis"
)

func main() {
	global.Init()
	defer global.DeInit()

	redisClient, err := redis.NewClient(config.RedisHost, config.RedisPassword)
	if err != nil {
		panic(err)
	}
	defer redis.Close(redisClient)

	activeEmailWorker := worker.NewSendActiveEmailWorker(redisClient)
	activeEmailWorker.Run()
}

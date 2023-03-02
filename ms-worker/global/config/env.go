package config

import "ms-workspace/package/env"

var (
	Environment = env.GetString("ENVIRONMENT", "development") // environment

	RedisHost     = env.GetString("REDIS_HOST", "redis:6379")        // redis host, include port
	RedisPassword = env.GetString("REDIS_PASSWORD", "redisPassword") // redis pw
	LogLevel      = env.GetInt64("LOG_LEVEL", -1)

	SendActiveEmailTopic    = env.GetString("SEND_ACTIVE_EMAIL_TOPIC", "send-active-email")
	NotificationServiceHost = env.GetString("NOTIFICATION_SERVICE_HOST", "ms-notification:82")
)

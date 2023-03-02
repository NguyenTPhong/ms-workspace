package config

import "ms-workspace/package/env"

var (
	Environment = env.GetString("ENVIRONMENT", "development") // environment

	MongoConnStr = env.GetString("MONGO_CONN_STR", "mongodb://mongo:27017")
	DBName       = env.GetString("DB_NAME", "ms-notification")

	RedisHost     = env.GetString("REDIS_HOST", "redis:6379")        // redis host, include port
	RedisPassword = env.GetString("REDIS_PASSWORD", "redisPassword") // redis pw
	LogLevel      = env.GetInt64("LOG_LEVEL", -1)

	SendGridAPIKey      = env.GetString("SENDGRID_API_KEY", "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")
	SendgridSenderEmail = env.GetString("SENDGRID_SENDER_EMAIL", "nguyenphong6102@gmail.com")
	SendgridSenderName  = env.GetString("SENDGRID_SENDER_EMAIL", "nguyenphong6102")
)

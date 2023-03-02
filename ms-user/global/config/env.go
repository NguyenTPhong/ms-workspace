package config

import "ms-workspace/package/env"

var (
	Environment = env.GetString("ENVIRONMENT", "development") // environment

	DbConnStr     = env.GetString("DB_CONN_STR", "host=postgres port=5432 user=dbUser password=dbPassword dbname=user_service sslmode=disable") // postgres connection string
	DbMaxConn     = env.GetInt64("DB_MAX_CONN", 10)                                                                                             // max connection to db
	DbMaxIdleConn = env.GetInt64("DB_MAX_IDLE_CONN", 2)                                                                                         // max idle connection to db
	DBLogLevel    = env.GetInt64("DB_LOG_LEVEL", 4)                                                                                             // db log level

	JWTKey        = env.GetString("JWT_KEY", "development-key")
	TokenLifeTime = env.GetInt64("TOKEN_LIFE_TIME", 60) // jwt token life time

	RedisHost     = env.GetString("REDIS_HOST", "redis:6379")        // redis host, include port
	RedisPassword = env.GetString("REDIS_PASSWORD", "redisPassword") // redis pw
	LogLevel      = env.GetInt64("LOG_LEVEL", -1)

	SendActiveEmailTopic = env.GetString("SEND_ACTIVE_EMAIL_TOPIC", "send-active-email")
)

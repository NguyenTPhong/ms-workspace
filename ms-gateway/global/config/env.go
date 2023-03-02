package config

import "ms-workspace/package/env"

var (
	Environment = env.GetString("ENVIRONMENT", "development") // environment
	Domain      = env.GetString("DOMAIN", "127.0.0.1")        // domain // app port
	LogLevel    = env.GetInt64("LOG_LEVEL", -1)

	UserServiceHost = env.GetString("USER_SERVICE_HOST", "ms-user:81")
)

package config

type HTTPServer struct {
}

type LogConfig struct {
}

type DBConfig struct {
}

type RedisConfig struct {
}

type JWTConfig struct {
}

type KeycloackConfig struct {
}

type MetricsConfig struct {
}

type RateLimitConfig struct {
}

type Config struct {
	HTTPServer HTTPServer
	Log        LogConfig
	DB         DBConfig
	Redis      RedisConfig
	JWT        JWTConfig
	Keycloack  KeycloackConfig
	Metrics    MetricsConfig
	RateLimit  RateLimitConfig
}

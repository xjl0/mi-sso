package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	DataBaseString string `env:"MI_DATABASE_STRING" env-default:"postgresql://admin:admin@localhost:5432/admin?sslmode=disable&application_name=golang"`
	JWTSign        string `env:"MI_JWT_SIGN" env-default:"secret"`
	JwtTTL         int    `env:"MI_JWT_TTL" env-default:"3600"`
	IsLocal        bool   `env:"MI_IS_LOCAL" env-default:"false"`
	LogLevel       string `env:"MI_LOG_LEVEL" env-default:"info"`
}

func MustLoad() Config {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	return cfg
}

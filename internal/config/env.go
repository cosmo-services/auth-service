package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Env struct {
	AppEnv string `mapstructure:"APP_ENV"`
	Port   string `mapstructure:"PORT"`

	PGHost string `mapstructure:"PG_HOST"`
	PGPort string `mapstructure:"PG_PORT"`
	PGUser string `mapstructure:"PG_USER"`
	PGPass string `mapstructure:"PG_PASS"`
	PGName string `mapstructure:"PG_NAME"`

	JwtSecreet    string        `mapstructure:"JWT_SECRET"`
	JwtAccessTTL  time.Duration `mapstructure:"JWT_ACCESS_TTL"`
	JwtRefreshTTL time.Duration `mapstructure:"JWT_REFRESH_TTL"`

	AllowedOrigins []string `mapstructure:"ALLOWED_ORIGINS"`
}

func NewEnv() Env {
	env := Env{}

	_, err := os.Stat(".env")
	useEnvFile := !os.IsNotExist(err)

	if useEnvFile {
		viper.SetConfigType("env")
		viper.SetConfigName(".env")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("Can't read the .env file: ", err)
		}

		err = viper.Unmarshal(&env)
		if err != nil {
			log.Fatal("Environment can't be loaded: ", err)
		}
	} else {
		env.bindEnv()
	}

	if env.AppEnv != "production" {
		log.Println("The App is running in development env")
	}

	return env
}

func (e *Env) bindEnv() {
	e.AppEnv = os.Getenv("APP_ENV")
	e.Port = os.Getenv("PORT")

	e.PGHost = os.Getenv("PG_HOST")
	e.PGPort = os.Getenv("PG_PORT")
	e.PGUser = os.Getenv("PG_USER")
	e.PGPass = os.Getenv("PG_PASS")
	e.PGName = os.Getenv("PG_NAME")

	e.JwtSecreet = os.Getenv("JWT_SECRET")

	if val := os.Getenv("JWT_ACCESS_TTL"); val != "" {
		d, err := time.ParseDuration(val)
		if err != nil {
			log.Fatalf("Invalid JWT_ACCESS_TTL format: %v", err)
		}
		e.JwtAccessTTL = d
	}

	if val := os.Getenv("JWT_REFRESH_TTL"); val != "" {
		d, err := time.ParseDuration(val)
		if err != nil {
			log.Fatalf("Invalid JWT_REFRESH_TTL format: %v", err)
		}
		e.JwtRefreshTTL = d
	}

	if val := os.Getenv("ALLOWED_ORIGINS"); val != "" {
		e.AllowedOrigins = strings.Split(val, ",")
	}
}

var Module = fx.Options(
	fx.Provide(NewEnv),
)

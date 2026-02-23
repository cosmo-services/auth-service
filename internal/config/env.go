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

	GmailPass string `mapstructure:"GMAIL_PASS"`
	GmailFrom string `mapstructure:"GMAIL_FROM"`
	GmailPort string `mapstructure:"GMAIL_PORT"`
	GmailSMTP string `mapstructure:"GMAIL_SMTP"`
	AppDomain string `mapstructure:"APP_DOMAIN"`

	NatsHost string `mapstructure:"NATS_HOST"`
	NatsPort string `mapstructure:"NATS_PORT"`
	NatsName string `mapstructure:"NATS_NAME"`
	NatsChan string `mapstructure:"NATS_CHAN"`

	JwtSecret     string        `mapstructure:"JWT_SECRET"`
	JwtAccessTTL  time.Duration `mapstructure:"JWT_ACCESS_TTL"`
	JwtRefreshTTL time.Duration `mapstructure:"JWT_REFRESH_TTL"`

	ApiActivationRoute string `mapstructure:"API_ACTIVATION_ROUTE"`

	TemplatesDir string `mapstructure:"TEMPLATES_DIR"`

	ClearExpiredTokensTTL  time.Duration `mapstructure:"CLEAR_EXP_TOKENS_TTL"`
	DeleteInactiveUsersTTL time.Duration `mapstructure:"DEL_INACTIVE_TTL"`

	MigrationPath string `mapstructure:"MIGRATION_PATH"`

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

	e.GmailPass = os.Getenv("GMAIL_PASS")
	e.GmailFrom = os.Getenv("GMAIL_FROM")
	e.GmailPort = os.Getenv("GMAIL_PORT")
	e.GmailSMTP = os.Getenv("GMAIL_SMTP")
	e.AppDomain = os.Getenv("APP_DOMAIN")

	e.NatsHost = os.Getenv("NATS_HOST")
	e.NatsPort = os.Getenv("NATS_PORT")
	e.NatsName = os.Getenv("NATS_NAME")
	e.NatsChan = os.Getenv("NATS_CHAN")

	e.JwtSecret = os.Getenv("JWT_SECRET")

	e.ApiActivationRoute = os.Getenv("API_ACTIVATION_ROUTE")

	e.MigrationPath = os.Getenv("MIGRATION_PATH")

	e.TemplatesDir = os.Getenv("TEMPLATES_DIR")

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

	if val := os.Getenv("CLEAR_EXP_TOKENS_TTL"); val != "" {
		d, err := time.ParseDuration(val)
		if err != nil {
			log.Fatalf("Invalid CLEAR_EXP_TOKENS_TTL format: %v", err)
		}
		e.ClearExpiredTokensTTL = d
	}

	if val := os.Getenv("DEL_INACTIVE_TTL"); val != "" {
		d, err := time.ParseDuration(val)
		if err != nil {
			log.Fatalf("Invalid DEL_INACTIVE_TTL format: %v", err)
		}
		e.DeleteInactiveUsersTTL = d
	}

	if val := os.Getenv("ALLOWED_ORIGINS"); val != "" {
		e.AllowedOrigins = strings.Split(val, ",")
	}
}

var Module = fx.Options(
	fx.Provide(NewEnv),
)

package initializers

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Configuration struct {
	Mode       string `mapstructure:"MODE"`
	ServerPort int    `mapstructure:"SERVER_POST"`
	MediaRoot  string `mapstructure:"MEDIA_ROOT"`

	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBPort     string `mapstructure:"POSTGRES_PORT"`
	DBUser     string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName     string `mapstructure:"POSTGRES_DB"`

	SnowflakeNode int64 `mapstructure:"SNOWFLAKE_NODE"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`

	EmailHost     string `mapstructure:"EMAIL_HOST"`
	EmailPort     int    `mapstructure:"EMAIL_PORT"`
	EmailUsername string `mapstructure:"EMAIL_USERNAME"`
	EmailPassword string `mapstructure:"EMAIL_PASSWORD"`
	EmailUseSSL   bool   `mapstructure:"EMAIL_USE_SSL"`
	EmailFrom     string `mapstructure:"EMAIL_FROM"`

	ThrottleByAnonymousIP uint `mapstructure:"THROTTLE_BY_ANONYMOUS_IP"`
}

var Config *Configuration

func InitConfig(path string) {
	if path == "" {
		path = os.Getenv("BASE_DIR")
		if path == "" {
			log.Panicln("please set config file path or set BASE_DIR")
		}
	}
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigType("env")
	v.SetConfigName("app")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Panicln("failed to load config file", err)
	}

	Config = &Configuration{}
	err = v.Unmarshal(Config)
	if err != nil {
		log.Panicln("failed to unmarshal config", err)
	}
}

package app

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPass      string `mapstructure:"DB_PASS"`
	DBName      string `mapstructure:"DB_NAME"`
	ChannelName string `mapstructure:"CHANNEL_NAME"`
	NatsServer  string `mapstructure:"NATS_SERVER"`
	NatsMonitor string `mapstructure:"NATS_MONITOR"`
	ServerHost  string `mapstructure:"SERVER_HOST"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}

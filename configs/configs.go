package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT           int    `mapstructure:"PORT"`
	ENV            string `mapstructure:"dev"`
	DB             string `mapstructure:"DB"`
	DB_HOST        string `mapstructure:"DB_HOST"`
	DB_USER        string `mapstructure:"DB_USER"`
	DB_PASSWORD    string `mapstructure:"DB_PASSWORD"`
	DB_NAME        string `mapstructure:"DB_NAME"`
	DB_PORT        string `mapstructure:"DB_PORT"`
	JWT_SECRET_KEY string `mapstructure:"JWT_SECRET_KEY"`
}

var Cfg Config

func LoadConfig(path string) (err error) {
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("ENV", "DEV")

	viper.SetDefault("DB", "postgres")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "admin")
	viper.SetDefault("DB_NAME", "BoardGame")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("JWT_SECRET_KEY", "secret")
	viper.AddConfigPath(path)
	viper.SetConfigName("DEV.env")
	viper.SetConfigType("env")
	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	viper.AutomaticEnv()

	if err = viper.Unmarshal(&Cfg); err != nil {
		return err
	}

	return nil
}

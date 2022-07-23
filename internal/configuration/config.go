package configuration

import (
	"fmt"
	"github.com/pozelim/go-hexagonal-example/internal/app/storage/postgres"
	"github.com/spf13/viper"
)

type Config struct {
	Port     int
	Name     string
	PathMap  string          `mapstructure:"path_map"`
	Postgres postgres.Config `mapstructure:"postgres"`
}

func New() Config {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}
	fmt.Println(viper.Get("appName"))
	return c
}

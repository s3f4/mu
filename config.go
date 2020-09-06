package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//config is main viper pointer
var config, env, test, development, production *viper.Viper

//setEnv assigns environment's viper instance to config
func setEnv() {
	switch env.Get("APP_ENV") {
	case "development":
		config = development
	case "test":
		config = test
	case "production":
		config = production
	default:
		config = development
	}
	config.Set("env", env.Get("APP_ENV"))
}

//createInstance initalizes all environment variables
func createInstance(file string) *viper.Viper {
	v := viper.New()

	if file != ".env" {
		v.SetConfigFile(fmt.Sprintf("./config/%v.json", file))
	} else {
		v.SetConfigFile(file)
	}

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return nil
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)

		if e.Name == ".env" {
			setEnv()
		}
	})

	return v
}

//New creates environments
func New() {
	env = createInstance(".env")
	development = createInstance("development")
	test = createInstance("test")
	production = createInstance("production")
	setEnv()
}

//GetConfig returns global config instance
func GetConfig() *viper.Viper {
	if config == nil {
		New()
	}
	return config
}

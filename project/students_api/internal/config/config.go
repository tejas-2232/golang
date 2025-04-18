package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address"`
}

// add HTTPServer struct to below struct
// struct embedding

type Config struct {
	//use struct annotations below for each type
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

// if there is error
// MustLoad function is returning the pointer of Config Struct

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to configuration file") // it's also checking the flags that are passed via cmd while running the go file
		flag.Parse()                                                     // parse all flags to configPath variable

		configPath = *flags // dereferenceing

		// the config path will either come from os.Getenv or Flag

		// if we still dont get config path then( if it's still empty)

		if configPath == "" {
			log.Fatal("Config path is not set")
		} //if

	} //if

	// check if there's a file availabel on path

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath) //Fatalf is used as we are formatting it with configPath
	}

	var cfg Config // var cfg - is of type is Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("can't read config file: %s", err.Error())
	}
	return &cfg // returning the address
}

package config

import (
	"flag"
	"log"
	"os"
)

type HTTPServer struct {
	Addr string
}

// add HTTPServer struct to below struct
// strcut embedding

type Config struct {
	//use struct annotations below for each type

	Env          string `yaml:"env" env:"ENV" env-required:"true"`
	storage_path string `yaml:"storage+path" env-required:"true"`

	HTTPServer `yaml:http_server`
}

// if there is error
func MustLoad() {
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

	// if there's a file avaiabel on path

	if _, err := os.Stat(configPath); os.IsNotExist(err) {

		log.Fatalf("Config file does not exist: %s", configPath)

	}

}

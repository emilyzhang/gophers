package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("hello world!")
	envFile := flag.String("config", "", "the .env configuration file to load")
	flag.Parse()
	fmt.Println(envFile)
	cfg, err := loadConfig(*envFile)
	fmt.Println(cfg)

	log := logrus.New()
	if err != nil {
		log.WithError(err).Fatal("Failed to load the configuration")
	}
}

type config struct {
	APIPort string

	DBHost string
	DBUser string 
	DBPassword string
	DBPort string
	DBName string
}


func loadConfig(file string) (config, error) {
	var cfg config
	if file != "" {
		if err := godotenv.Load(file); err != nil {
			return cfg, errors.Wrapf(err, "failed to load config from file: %s", file)
		}
	}
	if err := envconfig.Process("", &cfg); err != nil {
		return cfg, errors.Wrap(err, "failed to load the config from the env")
	}
	return cfg, nil
}

type api struct {
	logger logrus.FieldLogger
}

func (a *api) router(w http.ResponseWriter, req *http.Request) {
	a.logger.Printf("New request: %s", req.URL.Path)
	// pathPattern := regexp.MustCompile(`/(status|results)/\d+`)
	// fmt.Println(pathPattern)
	if req.URL.Path == "/crawl" && req.Method == http.MethodPost {
		return
	}
}

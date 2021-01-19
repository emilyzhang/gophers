package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		QuoteEmptyFields: true,
	})

	// simple loading of config
	envFile := flag.String("config", "", "the .env configuration file to load")
	flag.Parse()
	cfg, err := loadConfig(*envFile)
	if err != nil {
		log.WithError(err).Fatal("Failed to load the configuration")
	}
	log.Println(cfg)

	// start API
	a := api{logger: log, port: ":"+strconv.Itoa(cfg.APIPort)}
	a.start()
}

type config struct {
	APIPort int

	DBHost string
	DBUser string 
	DBPassword string
	DBPort int
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
	port string
}

func (a *api) router(w http.ResponseWriter, r *http.Request) {
	a.logger.Printf("New request: %s", r.URL.Path)
	if r.URL.Path == "/" {
		a.handleRoot(w, r)
		return
	}
}

// start starts the server.
func (a *api) start() {
	http.HandleFunc("/", a.router)
	a.logger.Print("Starting API server. Hello world!")
	http.ListenAndServe(a.port, nil)
	// TODO: implement graceful server shut down
}

func (a *api) handleRoot(w http.ResponseWriter, r *http.Request) {
	a.logger.Info("handling root")
}

package configs

import (
	"log"

	"github.com/timest/env"
)

const (
	// Dev develop environment
	Dev string = "dev"
	// Test test environment
	Test string = "test"
	// Pre preview environment
	Pre string = "pre"
	// Prod production environment
	Prod string = "prod"
)

//Env configuration in environment
var Env *config

type config struct {
	ProjectEnv string `env:"PROJECT_ENV" default:"dev"`
	APIVersion string `env:"API_VERSION" default:"Commit ID"`
	Mysql      struct {
		Host    string `default:"127.0.0.1"`
		Port    string `default:"3306"`
		DBName  string `default:"rank"`
		User    string `default:"root"`
		Pwd     string `default:"123"`
		Charset string `default:"utf8mb4"`
	}
	Redis struct {
		Host   string `default:"127.0.0.1"`
		Port   string `default:"6379"`
		Pwd    string `default:""`
		Prefix string `default:"rank|"`
	}
}

func init() {
	Env = new(config)
	env.IgnorePrefix()
	err := env.Fill(Env)
	log.Printf("configs: %+v\n", Env)
	if err != nil {
		panic(err)
	}
}

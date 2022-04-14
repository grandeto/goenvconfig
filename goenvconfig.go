package goenvconfig

import (
	"github.com/golang/glog"
	"github.com/joho/godotenv"
)

var Env string

func LoadConfig() {
	if Env == "" {
		panic("Env not set")
	}

	var err error

	switch Env {
    case "dev":
        err = godotenv.Load(".env.development")
    case "test":
        err = godotenv.Load(".env.test")
    case "prod":
        err = godotenv.Load()
	default:
		panic("Env is incorrect")
    }


	if err != nil {
		panic(err)
	}

	glog.Infof("%s environment config loaded successfully", Env)
}

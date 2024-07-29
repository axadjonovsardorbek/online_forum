package main

import (
	"api-gateway/api"
	cf "api-gateway/config"
	"api-gateway/config/logger"
	"log"
	"path/filepath"

	"api-gateway/grpc/client"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH) // Don't forget to change your log path

	services, err := client.NewGrpcClients(&config, *logger)
	if err != nil {
		log.Fatalf("error while connecting clients. err: %s", err.Error())
	}

	engine := api.NewRouter(services, *logger)

	err = engine.Run(config.HTTP_PORT)
	if err != nil {
		log.Fatalf("error while running server. err: %s", err.Error())
	}
}

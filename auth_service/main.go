package main

import (
	"auth-service/api"
	"auth-service/api/handlers"
	"auth-service/config"
	"auth-service/config/logger"
	"auth-service/postgresql"
	"auth-service/service"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	cf := config.Load()
	logger := logger.NewLogger(basepath, cf.LOG_PATH)
	em := config.NewErrorManager(logger)

	conn, err := postgresql.ConnectDB(&cf)
	em.CheckErr(err)
	defer conn.Close()

	us := service.NewUserService(conn)
	handler := handlers.NewHandler(us, *logger)

	roter := api.NewRouter(handler)
	logger.INFO.Println("Server is running on port ", cf.AUTH_PORT)
	if err := roter.Run(cf.AUTH_PORT); err != nil {
		logger.ERROR.Println(err)
	}
}

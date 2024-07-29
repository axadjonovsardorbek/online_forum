package handlers

import (
	"auth-service/config/logger"
	"auth-service/service"
)

type HTTPHandler struct {
	US     *service.UserService
	Logger logger.Logger
}

func NewHandler(us *service.UserService, l logger.Logger) *HTTPHandler {
	return &HTTPHandler{US: us, Logger: l}
}

package handlers

import (
	"api-gateway/grpc/client"
	"api-gateway/config/logger"
)

type Handler struct {
	srvs *client.GrpcClients
}

func NewHandler(srvs *client.GrpcClients, logger logger.Logger) *Handler {
	return &Handler{srvs: srvs}
}

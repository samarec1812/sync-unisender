package handler

import (
	"github.com/samarec1812/sync-unisender/auth/pkg/service"
	pb "github.com/samarec1812/sync-unisender/proto/auth"
)

type Handler struct {
	services *service.Service
	pb.AuthorizationAmoServer
	pb.AuthorizationUnisenderServer
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

package handler

import (
	"github.com/beanstalkd/go-beanstalk"
	"github.com/fasthttp/router"
	pb "github.com/samarec1812/sync-unisender/proto/auth"
	"github.com/samarec1812/sync-unisender/unisender/pkg/service"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	services  *service.Service
	clientRPC pb.AuthorizationUnisenderClient
	producer  *beanstalk.Conn
}

func NewHandler(services *service.Service, clientRPC pb.AuthorizationUnisenderClient, producer *beanstalk.Conn) *Handler {
	return &Handler{
		services:  services,
		clientRPC: clientRPC,
		producer:  producer,
	}
}

func (h *Handler) InitRoutes() func(ctx *fasthttp.RequestCtx) {

	r := router.New()
	r.POST("/token", h.saveToken)
	r.POST("/contacts", h.contactsMiddleware)

	return r.Handler
}

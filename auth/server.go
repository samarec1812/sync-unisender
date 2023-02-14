package auth

import (
	"github.com/samarec1812/sync-unisender/auth/pkg/handler"
	pb "github.com/samarec1812/sync-unisender/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewAuthServerGrpc(gserver *grpc.Server, handlers *handler.Handler) {
	pb.RegisterAuthorizationAmoServer(gserver, handlers)
	pb.RegisterAuthorizationUnisenderServer(gserver, handlers)
	reflection.Register(gserver)
}

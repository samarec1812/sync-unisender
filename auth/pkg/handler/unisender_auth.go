package handler

import (
	"context"

	"github.com/samarec1812/sync-unisender/auth/models"
	pb "github.com/samarec1812/sync-unisender/proto/auth"
)

func (h *Handler) SaveToken(ctx context.Context, in *pb.AuthUnisenderRequest) (*pb.AuthUnisenderResponse, error) {
	token := models.UnisenderUsers{
		Token:     in.Token,
		Uname:     in.Uname,
		AccountId: in.AccountId,
	}

	if err := token.Validate(); err != nil {
		return &pb.AuthUnisenderResponse{AuthCode: ""}, err
	}

	err := h.services.AuthorizationUnisender.SaveToken(token)
	if err != nil {
		return &pb.AuthUnisenderResponse{AuthCode: ""}, err
	}

	return &pb.AuthUnisenderResponse{AuthCode: "200"}, nil
}

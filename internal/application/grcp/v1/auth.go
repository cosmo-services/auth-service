package grpc_v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/cosmo-services/grpc-contracts/gen/api/v1"

	user_domain "main/internal/domain/user"
)

type AuthHandler struct {
	pb.UnimplementedSocialServiceServer

	userService *user_domain.UserService
}

func NewAuthHandler(
	userService *user_domain.UserService,
) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

func (h *AuthHandler) GetUser(
	ctx context.Context,
	req *pb.GetUserRequest,
) (*pb.GetUserResponse, error) {
	if req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	p, err := h.userService.GetUser(req.UserId)
	if err != nil {
		if errors.Is(err, user_domain.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetUserResponse{
		User: &pb.UserData{
			Id:        p.ID,
			Username:  p.Username,
			Email:     p.Email,
			IsActive:  p.IsActive,
			CreatedAt: p.CreatedAt.Unix(),
		},
	}, nil
}

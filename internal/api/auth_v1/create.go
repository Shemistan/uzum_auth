package auth_v1

import (
	"context"
	"github.com/Shemistan/uzum_auth/internal/models"
	pb "github.com/Shemistan/uzum_auth/pkg/auth_v1"
)

func (a *Auth) Create(ctx context.Context, req *pb.Create_Request) (*pb.Create_Response, error) {
	defer recoveryFunction(ctx)
	res, err := a.AuthService.Create(ctx, a.getModelUser(req.User))
	if err != nil {
		return nil, err
	}

	return &pb.Create_Response{AccessToken: res.Access}, err
}

func (a *Auth) getModelUser(req *pb.User) *models.User {
	return &models.User{
		Login:      req.Login,
		Password:   req.Password,
		Role:       req.Role,
		DateOfBirt: req.DateOfBirt.AsTime(),
	}
}

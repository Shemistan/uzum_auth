package auth_v1

import (
	"context"

	"github.com/Shemistan/uzum_auth/internal/models"
	pb "github.com/Shemistan/uzum_auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *Auth) GetAllUsers(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllUsers_Response, error) {
	res, err := a.AuthService.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*pb.User, 0, len(res))

	for _, v := range res {
		users = append(users, getPbUser(v))
	}

	return &pb.GetAllUsers_Response{
		User: users,
	}, err
}

func getPbUser(req *models.User) *pb.User {
	return &pb.User{
		Login:    req.Login,
		Password: req.Password,
		Role:     req.Role,
		//DateOfBirt: timestamppb.Timestamp{} req.DateOfBirt.,
	}
}

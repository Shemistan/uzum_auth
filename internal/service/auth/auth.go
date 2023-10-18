package auth

import (
	"context"

	"github.com/Shemistan/uzum_auth/internal/models"
	repo "github.com/Shemistan/uzum_auth/internal/storage"
)

type IAuthSystemService interface {
	Create(ctx context.Context, req *models.User) (*models.Token, error)
	Update(ctx context.Context, req *models.User) error
	Auth(ctx context.Context, req *models.User) (*models.Token, error)
	Delete(ctx context.Context, req uint32) error
	CheckToken(ctx context.Context, req *models.Token) error

	GetUser(ctx context.Context, req uint32) (*models.User, error)
	GetUsers(ctx context.Context, req []uint32) ([]*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
}

type authSystemService struct {
	appConfig *models.Config
	storage   repo.IStorage
}

func NewAuthSystemService(appConfig *models.Config, storage repo.IStorage) IAuthSystemService {
	return &authSystemService{
		appConfig: appConfig,
		storage:   storage,
	}
}

func (a authSystemService) Create(ctx context.Context, req *models.User) (*models.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (a authSystemService) Update(ctx context.Context, req *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (a authSystemService) Auth(ctx context.Context, req *models.User) (*models.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (a authSystemService) Delete(ctx context.Context, req uint32) error {
	//TODO implement me
	panic("implement me")
}

func (a authSystemService) CheckToken(ctx context.Context, req *models.Token) error {
	//TODO implement me
	panic("implement me")
}

func (a authSystemService) GetUser(ctx context.Context, req uint32) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a authSystemService) GetUsers(ctx context.Context, req []uint32) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a authSystemService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

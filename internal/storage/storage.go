package storage

import (
	"context"
	"github.com/Shemistan/uzum_auth/internal/models"
)

//go:generate mockgen -destination=mocks/mock_repo.go -package=mocks . Repo

type IStorage interface {
	GetUser(ctx context.Context, id int64) *models.User
}

package postgresql

import (
	"github.com/jmoiron/sqlx"

	repo "github.com/Shemistan/uzum_auth/internal/storage"
)

type storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) repo.IStorage {
	return &storage{db: db}
}

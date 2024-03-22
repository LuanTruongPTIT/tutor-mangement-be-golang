package repository

import (
	"context"

	"github.com/LuanTruongPTIT/tutor-be/modules/auth/model"
	"github.com/LuanTruongPTIT/tutor-be/pkg/dbs"
)

type IAuthRepository interface {
	Create(ctx context.Context, auth *model.Auth) error
}

type AuthRepo struct {
	db dbs.IDatabase
}

func NewAuthRepository(db dbs.IDatabase) *AuthRepo {
	return &AuthRepo{db: db}
}
func (r *AuthRepo) Create(ctx context.Context, auth *model.Auth) error {
	return r.db.Create(ctx, auth)
}

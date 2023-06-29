package store

import (
	"context"

	"github.com/Malarkey-Jhu/miniblog/internal/pkg/model"
	"gorm.io/gorm"
)

type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
}

var _ UserStore = (*users)(nil)

type users struct {
	db *gorm.DB
}

func newUsers(db *gorm.DB) *users {
	return &users{db}
}

func (us *users) Create(ctx context.Context, user *model.UserM) error {
	return us.db.Create(user).Error
}

package user

import (
	"context"
	"regexp"

	"github.com/Malarkey-Jhu/miniblog/internal/miniblog/store"
	"github.com/Malarkey-Jhu/miniblog/internal/pkg/errno"
	"github.com/Malarkey-Jhu/miniblog/internal/pkg/model"
	v1 "github.com/Malarkey-Jhu/miniblog/pkg/api/miniblog/v1"
	"github.com/jinzhu/copier"
)

type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
}

type userBiz struct {
	ds store.IStore
}

var _ UserBiz = (*userBiz)(nil)

func New(ds store.IStore) UserBiz {
	return &userBiz{
		ds: ds,
	}
}

func (u *userBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, r)

	if err := u.ds.Users().Create(ctx, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}

	return nil
}

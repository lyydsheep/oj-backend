package service

import (
	"back/api/reply"
	"back/api/request"
	"back/common/errcode"
	"back/common/util"
	"back/logic/domain"
	"back/logic/repository"
	"context"
)

type UserServiceV1 struct {
	repo repository.UserRepository
}

func (u *UserServiceV1) Login(ctx context.Context, user *request.User) error {
	domainUser, err := u.repo.GetUserByAccount(ctx, user.UserAccount)
	if err != nil {
		return err
	}
	// TODO 对密码进行加密处理
	if domainUser.UserPassword != user.UserPassword {
		return errcode.ErrNotFound.WithCause(err).AppendMsg("password is wrong")
	}
	return nil
}

func (u *UserServiceV1) Register(ctx context.Context, user *request.User) error {
	var domainUser domain.User
	if err := util.Convert(user, &domainUser); err != nil {
		return errcode.ErrServer.WithCause(err).AppendMsg("fail to convert request.User to domainUser.User")
	}
	// TODO 需要检查账号是否已经存在
	// TODO 需要对密码进行加密存储
	return u.repo.StoreUser(ctx, &domainUser)
}

func (u *UserServiceV1) GetUserByUid(ctx context.Context, uid string) (*reply.User, error) {
	domainUser, err := u.repo.GetUserByUid(ctx, uid)
	if err != nil {
		return nil, err
	}
	var user reply.User
	if err = util.Convert(domainUser, &user); err != nil {
		return nil, errcode.ErrServer.WithCause(err).AppendMsg("fail to convert domainUser.User to reply.User")
	}
	return &user, nil
}

func NewUserServiceV1(repo repository.UserRepository) UserService {
	return &UserServiceV1{repo: repo}
}

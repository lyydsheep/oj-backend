package repository

import (
	"back/common/errcode"
	"back/common/util"
	"back/dal/dao"
	"back/dal/model"
	"back/logic/domain"
	"context"
)

type UserRepositoryV1 struct {
	Query *dao.Query
}

func (u *UserRepositoryV1) GetUserByAccount(ctx context.Context, account string) (*domain.User, error) {
	entity, err := u.Query.User.WithContext(ctx).Where(u.Query.User.UserAccount.Eq(account)).First()
	if err != nil {
		return nil, errcode.ErrNotFound.WithCause(err).AppendMsg("not find user, account: " + account)
	}
	var user domain.User
	if err = util.Convert(entity, &user); err != nil {
		return nil, errcode.ErrServer.WithCause(err).AppendMsg("convert entity to domain.User error")
	}
	return &user, nil
}

func NewUserRepositoryV1(query *dao.Query) UserRepository {
	return &UserRepositoryV1{
		Query: query,
	}
}

func (u *UserRepositoryV1) GetUserByUid(ctx context.Context, uid string) (*domain.User, error) {
	entity, err := u.Query.User.WithContext(ctx).Where(u.Query.User.UID.Eq(uid)).First()
	if err != nil {
		return nil, errcode.ErrNotFound.WithCause(err).AppendMsg("not find user, uid: " + uid)
	}
	var user domain.User
	if err = util.Convert(entity, &user); err != nil {
		return nil, errcode.ErrServer.WithCause(err).AppendMsg("convert entity to domain.User error")
	}
	return &user, nil
}

func (u *UserRepositoryV1) StoreUser(ctx context.Context, user *domain.User) error {
	var entity model.User
	if err := util.Convert(user, &entity); err != nil {
		return errcode.ErrServer.WithCause(err).AppendMsg("convert domain.User to entity error")
	}
	err := u.Query.User.WithContext(ctx).Create(&entity)
	if err != nil {
		return errcode.ErrServer.WithCause(err).AppendMsg("create user error")
	}
	return nil
}

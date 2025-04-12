package service

import (
	"back/api/reply"
	"back/api/request"
	"context"
)

type QuestionService interface {
	CreateQuestion(ctx context.Context, question *request.CreateQuestion) (int64, error)
	GetQuestionById(ctx context.Context, id int64) (*reply.Question, error)
}

type UserService interface {
	Register(ctx context.Context, user *request.User) error
	GetUserByUid(ctx context.Context, uid string) (*reply.User, error)
	Login(ctx context.Context, user *request.User) error
}

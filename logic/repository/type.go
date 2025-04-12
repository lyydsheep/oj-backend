package repository

import (
	"back/logic/domain"
	"context"
)

type UserRepository interface {
	GetUserByUid(context.Context, string) (*domain.User, error)
	StoreUser(context.Context, *domain.User) error
	GetUserByAccount(context.Context, string) (*domain.User, error)
}

type QuestionRepository interface {
	GetQuestionById(context.Context, int64) (*domain.Question, error)
	StoreQuestion(context.Context, *domain.Question) (int64, error)
}

type QuestionSubmitRepository interface {
	CreateQuestionSubmit(context.Context, *domain.QuestionSubmit) (*domain.QuestionSubmit, error)
	GetQuestionSubmitById(context.Context, int64) (*domain.QuestionSubmit, error)
	GetQuestionSubmitListByUserId(context.Context, int64) ([]domain.QuestionSubmit, error)
	GetQuestionSubmitListByQuestionId(context.Context, int64) ([]domain.QuestionSubmit, error)
	GetQuestionSubmitListByQuestionIdAndUserId(context.Context, int64, int64) ([]domain.QuestionSubmit, error)
}

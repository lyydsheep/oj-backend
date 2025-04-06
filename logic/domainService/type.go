package domainService

import (
	"back/logic/domain"
	"context"
)

type UserDomainService interface {
	GetUser(context.Context, string) (*domain.User, error)
	CreateUser(context.Context, *domain.User) (*domain.User, error)
}

type QuestionDomainService interface {
	GetQuestionById(context.Context, int64) (*domain.Question, error)
	StoreQuestion(context.Context, *domain.Question) (int64, error)
}

type QuestionSubmitDomainService interface {
	CreateQuestionSubmit(context.Context, *domain.QuestionSubmit) (*domain.QuestionSubmit, error)
	GetQuestionSubmitById(context.Context, int64) (*domain.QuestionSubmit, error)
	GetQuestionSubmitListByUserId(context.Context, int64) ([]domain.QuestionSubmit, error)
	GetQuestionSubmitListByQuestionId(context.Context, int64) ([]domain.QuestionSubmit, error)
	GetQuestionSubmitListByQuestionIdAndUserId(context.Context, int64, int64) ([]domain.QuestionSubmit, error)
}

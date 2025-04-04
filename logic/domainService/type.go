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
	GetQuestionByTitle(context.Context, string) (*domain.Question, error)
	GetQuestionById(context.Context, int64) (*domain.Question, error)
	CreateQuestion(context.Context, *domain.Question) (*domain.Question, error)
	EditQuestion(context.Context, *domain.Question) (*domain.Question, error)
	// DeleteQuestion 根据id删除题目
	DeleteQuestion(context.Context, int64) (bool, error)
}

type QuestionSubmitDomainService interface {
	CreateQuestionSubmit(context.Context, *domain.QuestionSubmit) (*domain.QuestionSubmit, error)
	GetQuestionSubmitById(context.Context, int64) (*domain.QuestionSubmit, error)
	GetQuestionSubmitListByUserId(context.Context, int64) ([]domain.QuestionSubmit, error)
	GetQuestionSubmitListByQuestionId(context.Context, int64) ([]domain.QuestionSubmit, error)
	GetQuestionSubmitListByQuestionIdAndUserId(context.Context, int64, int64) ([]domain.QuestionSubmit, error)
}

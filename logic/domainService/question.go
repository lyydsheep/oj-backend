package domainService

import (
	"back/common/errcode"
	"back/common/util"
	"back/dal/dao"
	"back/dal/model"
	"back/logic/domain"
	"context"
	"github.com/jinzhu/copier"
)

type QuestionDomainServiceV1 struct {
	query dao.Query
}

func NewQuestionDomainServiceV1(query dao.Query) *QuestionDomainServiceV1 {
	return &QuestionDomainServiceV1{
		query: query,
	}
}

func (q *QuestionDomainServiceV1) GetQuestionByTitle(ctx context.Context, s string) (*domain.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (q *QuestionDomainServiceV1) GetQuestionById(ctx context.Context, i int64) (*domain.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (q *QuestionDomainServiceV1) CreateQuestion(ctx context.Context, question *domain.Question) (*domain.Question, error) {
	var q model.Question
	err := copier.Copy(&q, question)
	if err != nil {
		return nil, errcode.Wrap("copy error", err)
	}
}

func (q *QuestionDomainServiceV1) EditQuestion(ctx context.Context, question *domain.Question) (*domain.Question, error) {
	//TODO implement me
	panic("implement me")
}

func (q *QuestionDomainServiceV1) DeleteQuestion(ctx context.Context, i int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

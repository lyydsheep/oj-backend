package domainService

import (
	"back/common/errcode"
	"back/dal/dao"
	"back/dal/model"
	"back/logic/domain"
	"context"
	"github.com/jinzhu/copier"
)

type QuestionDomainServiceV1 struct {
	query *dao.Query
}

func NewQuestionDomainServiceV1(query *dao.Query) QuestionDomainService {
	return &QuestionDomainServiceV1{
		query: query,
	}
}

func (q *QuestionDomainServiceV1) GetQuestionById(ctx context.Context, i int64) (*domain.Question, error) {
	entity, err := q.query.Question.WithContext(ctx).Where(q.query.Question.ID.Eq(i)).First()
	if err != nil {
		return nil, errcode.ErrServer.WithCause(err).AppendMsg("query entity error")
	}
	var question domain.Question
	if err = copier.Copy(&question, entity); err != nil {
		return nil, errcode.ErrServer.WithCause(err).AppendMsg("entity to question error")
	}
	return &question, nil
}

func (q *QuestionDomainServiceV1) StoreQuestion(ctx context.Context, question *domain.Question) (int64, error) {
	var questionModel model.Question
	if err := copier.Copy(&questionModel, question); err != nil {
		return 0, errcode.ErrServer.WithCause(err).AppendMsg("domain to entity error")
	}
	if err := q.query.Question.WithContext(ctx).Create(&questionModel); err != nil {
		return 0, errcode.ErrServer.WithCause(err).AppendMsg("create entity error")
	}
	return questionModel.ID, nil
}

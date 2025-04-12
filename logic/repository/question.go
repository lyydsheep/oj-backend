package repository

import (
	"back/common/errcode"
	"back/common/util"
	"back/dal/dao"
	"back/dal/model"
	"back/logic/domain"
	"context"
	"fmt"
)

type QuestionRepositoryV1 struct {
	query *dao.Query
}

func NewQuestionRepositoryV1(query *dao.Query) QuestionRepository {
	return &QuestionRepositoryV1{
		query: query,
	}
}

func (q *QuestionRepositoryV1) GetQuestionById(ctx context.Context, i int64) (*domain.Question, error) {
	// 如果没有找到，会报错
	entity, err := q.query.Question.WithContext(ctx).Where(q.query.Question.ID.Eq(i)).First()
	if err != nil {
		return nil, errcode.ErrServer.WithCause(err).AppendMsg(fmt.Sprintf("query entity error, err:%s, question_id:%d", err.Error(), i))
	}
	var question domain.Question
	if err = util.Convert(entity, &question); err != nil {
		return nil, errcode.ErrServer.WithCause(err).AppendMsg("entity to question error")
	}
	return &question, nil
}

func (q *QuestionRepositoryV1) StoreQuestion(ctx context.Context, question *domain.Question) (int64, error) {
	var questionModel model.Question
	if err := util.Convert(question, &questionModel); err != nil {
		return 0, errcode.ErrServer.WithCause(err).AppendMsg("domain to entity error")
	}
	if err := q.query.Question.WithContext(ctx).Create(&questionModel); err != nil {
		return 0, errcode.ErrServer.WithCause(err).AppendMsg("create entity error")
	}
	return questionModel.ID, nil
}

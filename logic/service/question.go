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

type QuestionServiceV1 struct {
	svc repository.QuestionRepository
}

// GetQuestionById implements QuestionService.
func (q *QuestionServiceV1) GetQuestionById(ctx context.Context, id int64) (*reply.Question, error) {
	domainQuestion, err := q.svc.GetQuestionById(ctx, id)
	if err != nil {
		return nil, err
	}
	var question reply.Question
	if err := util.Convert(domainQuestion, &question); err != nil {
		return nil, errcode.ErrServer.WithCause(err).AppendMsg("copy question error")
	}
	return &question, nil
}

func NewQuestionServiceV1(svc repository.QuestionRepository) QuestionService {
	return &QuestionServiceV1{svc: svc}
}

func (q *QuestionServiceV1) CreateQuestion(ctx context.Context, question *request.CreateQuestion) (int64, error) {
	var domainQuestion domain.Question
	if err := util.Convert(question, &domainQuestion); err != nil {
		return 0, errcode.ErrServer.WithCause(err).AppendMsg("copy question error")
	}
	// TODO 进行鉴权
	// TODO 补充缺省值
	return q.svc.StoreQuestion(ctx, &domainQuestion)
}

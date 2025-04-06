package appService

import (
	"back/api/request"
	"back/common/errcode"
	"back/logic/domain"
	"back/logic/domainService"
	"context"
	"github.com/jinzhu/copier"
)

type QuestionAppServiceV1 struct {
	svc domainService.QuestionDomainService
}

func NewQuestionAppServiceV1(svc domainService.QuestionDomainService) QuestionAppService {
	return &QuestionAppServiceV1{svc: svc}
}

func (q QuestionAppServiceV1) CreateQuestion(ctx context.Context, question *request.CreateRequest) (int64, error) {
	var domainQuestion domain.Question
	if err := copier.Copy(&domainQuestion, question); err != nil {
		return 0, errcode.ErrServer.WithCause(err).AppendMsg("copy question error")
	}
	// TODO 进行鉴权
	// TODO 补充缺省值
	return q.svc.StoreQuestion(ctx, &domainQuestion)
}

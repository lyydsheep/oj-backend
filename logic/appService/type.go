package appService

import (
	"back/api/request"
	"context"
)

type QuestionAppService interface {
	CreateQuestion(ctx context.Context, question *request.CreateRequest) (int64, error)
}

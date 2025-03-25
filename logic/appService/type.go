package appService

import (
	"context"
)

type DemoAppService interface {
	GetAllIdentities(c context.Context) ([]int64, error)
	CreateDemoOrder(c context.Context, order *request.DemoOrderReq) (*reply.DemoOrder, error)
}

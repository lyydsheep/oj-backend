package dao

import (
	"context"
)

type DemoDAO interface {
	FindAllDemo(c context.Context) ([]model.DemoOrder, error)
	CreateDemoOrder(c context.Context, order *domain.DemoOrder) (*model.DemoOrder, error)
}

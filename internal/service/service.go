package service

import (
	"context"
	"goproject/main/ginweb/internal/dao"
)

type Service struct {
	Ctx context.Context
	Dao *dao.Dao
}

func New(ctx context.Context) Service {
	return Service{
		Ctx: ctx,
		Dao: dao.New(),
	}
}

package controller

import (
	"context"
	v1 "zzw/api/v1"
	"zzw/internal/model"
	"zzw/internal/service"
)

var (
	Index = cIndex{}
)

type cIndex struct {
}

func (*cIndex) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	service.View().Render(ctx, model.View{
		Title: "index",
	})
	return
}

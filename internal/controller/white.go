package controller

import (
	"context"
	"log"
	v1 "zzw/api/v1"
	"zzw/internal/service"
)

var (
	White = cWhite{}
)

type cWhite struct {
}

func (c *cWhite) White(ctx context.Context, req *v1.WReq) (res *v1.WRes, err error) {
	res = &v1.WRes{}
	iswhite, err := service.White().IsWhite(ctx, req.Address)
	if err != nil {

	}
	res.IsWhite = iswhite
	return
}

func (c *cWhite) AddWhite(ctx context.Context, req *v1.AddWReq) (res *v1.AddWRes, err error) {
	res1 := &v1.AddWRes{}
	err1 := service.White().AddWhite(ctx, req.AddressList)
	if err1 != nil {
		log.Println(err1)
		return
	}
	res1.Result = "add success"
	return
}

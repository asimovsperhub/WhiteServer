package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"log"
	"strings"
	"zzw/internal/dao"
)

func White() *sWhite {
	return &sWhite{}
}

type sWhite struct {
}

func (s *sWhite) IsWhite(ctx context.Context, query string) (bool, error) {
	n, err := dao.Address.Ctx(ctx).Where(dao.Address.Columns().Addr, query).Count()
	if err != nil {
		log.Println(err)
		return false, err
	}
	if n > 0 {
		return true, nil
	} else {
		return false, nil
	}

}

func (s *sWhite) AddWhite(ctx context.Context, WhiteList string) error {
	wl := strings.ReplaceAll(WhiteList, "\n", " ")
	whitelist := strings.Split(wl, " ")
	log.Println(whitelist, len(whitelist))
	for _, addr := range whitelist {
		n, err := dao.Address.Ctx(ctx).Where(dao.Address.Columns().Addr, addr).Count()
		if err != nil {
			log.Println(err)
			return err
		}
		if n < 1 {
			res, err1 := dao.Address.Ctx(ctx).Insert(g.Map{"addr": addr})
			if err1 != nil {
				log.Println("Insert failed", err1)
				return err1
			}
			log.Println(res)
		}
	}
	return nil
}

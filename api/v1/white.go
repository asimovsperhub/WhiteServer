package v1

import "github.com/gogf/gf/v2/frame/g"

type WReq struct {
	g.Meta  `path:"/api/isWhite" tags:"Hello" method:"get" `
	Address string `json:"key" v:"required" dc:"Address"`
}
type WRes struct {
	g.Meta  `mime:"text/html" example:"string"`
	IsWhite bool `json:"is_white"`
}

type AddWReq struct {
	g.Meta      `path:"/api/addWhite" tags:"Hello" method:"post" `
	AddressList string `json:"address_list" v:"required" dc:"address_list"`
}
type AddWRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Result string `json:"result"`
}

// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Address is the golang structure of table Address for DAO operations like Where/Data.
type Address struct {
	g.Meta  `orm:"table:Address, do:true"`
	Id      interface{} // UID
	Addr    interface{} // 账户地址
	Iswhite interface{} // 是否白名单
}
package mysql

import (
	"ts-error/err_code"
)

type MysqlOpType err_code.ErrCodeSub
type MysqlOpObj err_code.ErrCodeSubSub

const (
	Insert MysqlOpType = iota + 1
	Delete
	Update
	Search
	Other
)

var MysqlOpTypeMsg = map[MysqlOpType]string{
	Insert: "Insert",
	Delete: "Delete",
	Update: "Update",
	Search: "Search",
	Other:  "Other",
}

const (
	Order MysqlOpObj = iota + 1
	Symbol
	Source
)

var MysqlOpObjMsg = map[MysqlOpObj]string{
	Order:  "Order",
	Symbol: "Symbol",
	Source: "Source",
}


package args

import "ts-error/err_code"

type ArgsType err_code.ErrCodeSub

const (
	Trade ArgsType = iota + 1
)

var ArgsTypeMsg = map[ArgsType]string{
	Trade: "Trade",
}
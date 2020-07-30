package system

import "ts-error/err_code"

type timeoutType err_code.ErrCodeSubSub

const (
	Trade timeoutType = iota + 1
)

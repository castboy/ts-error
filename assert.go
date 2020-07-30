package error

import "ts-error/err_code"

var errMsgAssert = map[err_code.ErrCodeSub]string{}

func AssertErrorSyncOrder(i ...interface{}) {
	// TODO
}

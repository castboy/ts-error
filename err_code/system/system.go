package system

import "ts-error/err_code"

type Type err_code.ErrCodeSub

const (
	SystemErr Type = iota + 1 // package any panic.
	SystemBusy // to reject user request.
	JobTimeout
	ClientFailed // client connect failed
)

var SystemTypeMsg = map[Type]string {
	SystemErr:    "System Error",
	SystemBusy:   "System Busy",
	JobTimeout:   "Job Timeout",
	ClientFailed: "Client Conn Failed",
}

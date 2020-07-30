package error

import (
	"ts-error/err_code"
	"ts-error/err_code/system"
)

type SystemErrer struct {
	*baseErrer
}

func NewSystemErrer(errOrigin error, Type system.Type) *SystemErrer {
	base := newErr()
	base.init(errOrigin, err_code.SystemErr, err_code.ErrCodeSub(Type), err_code.ErrCodeSubSub(0))

	return &SystemErrer{base}
}

func (i *SystemErrer) getErrCodeSubMsg() string {
	return system.SystemTypeMsg[system.Type(i.codeSub)]
}

func (i *SystemErrer) getErrCodeSubSubMsg() string {
	return ""
}

func (i *SystemErrer) Error() string {
	return encodeErrMsg(i)
}


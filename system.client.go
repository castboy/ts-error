package error

import (
	"ts-error/err_code"
	"ts-error/err_code/system"
)

type SystemClientErrer struct {
	*baseErrer
}

func NewSystemClientErrer(errOrigin error, Type system.Type, client system.ClientType) *SystemClientErrer {
	base := newErr()
	base.init(errOrigin, err_code.SystemErr, err_code.ErrCodeSub(Type), err_code.ErrCodeSubSub(client))

	return &SystemClientErrer{base}
}

func (i *SystemClientErrer) getErrCodeSubMsg() string {
	return system.SystemTypeMsg[system.Type(i.codeSub)]
}

func (i *SystemClientErrer) getErrCodeSubSubMsg() string {
	return system.ClientTypeMsg[system.ClientType(i.codeSubSub)]
}

func (i *SystemClientErrer) Error() string {
	return encodeErrMsg(i)
}

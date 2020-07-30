package error

import (
	"ts-error/err_code"
	"ts-error/err_code/mysql"
)

type MysqlErrer struct {
	*baseErrer
}

func NewMysqlErrer(errOrigin error, opType mysql.MysqlOpType, opObj mysql.MysqlOpObj, comment ...Comment) *MysqlErrer {
	base := newErr()
	base.init(errOrigin, err_code.MysqlErr, err_code.ErrCodeSub(opType), err_code.ErrCodeSubSub(opObj), comment...)

	return &MysqlErrer{base}
}

func (i *MysqlErrer) getErrCodeSubMsg() string {
	return mysql.MysqlOpTypeMsg[mysql.MysqlOpType(i.codeSub)]
}

func (i *MysqlErrer) getErrCodeSubSubMsg() string {
	return mysql.MysqlOpObjMsg[mysql.MysqlOpObj(i.codeSubSub)]
}

func (i *MysqlErrer) Error() string {
	return encodeErrMsg(i)
}

package error

import (
	"ts-error/err_code"
	"ts-error/err_code/kafka"
)

type KafkaErrer struct {
	*baseErrer
}

func NewKafkaErrer(errOrigin error, opType kafka.KafkaOpType, opObj kafka.KafkaOpObj, comment ...Comment) *KafkaErrer {
	base := newErr()
	base.init(errOrigin, err_code.MysqlErr, err_code.ErrCodeSub(opType), err_code.ErrCodeSubSub(opObj), comment...)

	return &KafkaErrer{base}
}

func (i *KafkaErrer) getErrCodeSubMsg() string {
	return kafka.KafkaOpTypeMsg[kafka.KafkaOpType(i.codeSub)]
}

func (i *KafkaErrer) getErrCodeSubSubMsg() string {
	return kafka.KafkaOpObjMsg[kafka.KafkaOpObj(i.codeSubSub)]
}

func (i *KafkaErrer) Error() string {
	return encodeErrMsg(i)
}
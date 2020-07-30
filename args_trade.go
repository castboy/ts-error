package error

import (
	"ts-error/err_code"
	"ts-error/err_code/args"
)

type TradeArgsErrer struct {
	*baseErrer
}

func NewTradeArgsErrer(errOrigin error, argsType args.ArgsType, tradeArg args.TradeArg, comment ...Comment) *TradeArgsErrer {
	base := newErr()
	base.init(errOrigin, err_code.MysqlErr, err_code.ErrCodeSub(argsType), err_code.ErrCodeSubSub(tradeArg), comment...)

	return &TradeArgsErrer{base}
}

func (i *TradeArgsErrer) getErrCodeSubMsg() string {
	return args.ArgsTypeMsg[args.ArgsType(i.codeSub)]
}

func (i *TradeArgsErrer) getErrCodeSubSubMsg() string {
	return args.TradeArgsMsg[args.TradeArg(i.codeSubSub)]
}

func (i *TradeArgsErrer) Error() string {
	return encodeErrMsg(i)
}
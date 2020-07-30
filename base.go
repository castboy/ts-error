package error

import (
	"encoding/json"
	"ts-error/err_code"
)

type baseErrer struct {
	code       err_code.ErrCode
	codeSub    err_code.ErrCodeSub
	codeSubSub err_code.ErrCodeSubSub
	originErr  error
	callFunc   []string
	comment    []Comment
}

type Comment struct {
	k string
	v interface{}
}

func newErr() *baseErrer {
	res := &baseErrer{}
	return res
}

func (me *baseErrer) init(errOrigin error, code err_code.ErrCode, codeSub err_code.ErrCodeSub, codeSubSub err_code.ErrCodeSubSub, comment ...Comment) {
	me.setCode(code)
	me.setCodeSub(codeSub)
	me.setCodeSubSub(codeSubSub)
	me.setOriginErr(errOrigin)

	f := getCallFunc(newDeep)
	me.appendCallFunc(f)

	me.appendComment(comment...)
}

func (me *baseErrer) setCode(code err_code.ErrCode) {
	me.code = code
}

func (me *baseErrer) setCodeSub(sub err_code.ErrCodeSub) {
	me.codeSub = sub
}

func (me *baseErrer) setCodeSubSub(subSub err_code.ErrCodeSubSub) {
	me.codeSubSub = subSub
}

func (me *baseErrer) setOriginErr(err error) {
	me.originErr = err
}

func (me *baseErrer) appendCallFunc(f ...string) {
	me.callFunc = append(me.callFunc, f...)
}

func (me *baseErrer) appendComment(cmt ...Comment) {
	me.comment = append(me.comment, cmt...)
}

func (me *baseErrer) encodeCallFunc() string {
	l := len(me.callFunc)
	s := ""
	for i := l-1; i >= 0; i-- {
		s += me.callFunc[i]
		if i != 0 {
			s += "->"
		}
	}

	return s
}

func (me *baseErrer) encodeComment() string {
	comment, _ := json.Marshal(me.comment)
	return string(comment)
}

func (me *baseErrer) getOriginErr() error {
	return me.originErr
}

func (i *baseErrer) getErrCodeMsg() string {
	return err_code.ErrCodeMsg[i.code]
}
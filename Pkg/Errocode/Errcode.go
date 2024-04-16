package Errocode

import "fmt"

type Error struct {
	code string
	msg  string
}

var _codes = map[string]string{}

func NewError(code string, msg string) *Error {
	if _, ok := _codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	_codes[code] = msg
	return &Error{code: code, msg: msg}
}
func (e *Error) Eeror() string {
	return fmt.Sprintf("错误码：%s, 错误信息：%s", e.Code(), e.Msg())
}
func (e *Error) Code() string {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

package werror

import "github.com/gookit/goutil/errorx"

type BizError struct {
	errorx.ErrorX
	Code int
}

func NewBizError(msg string, code ...int) *BizError {
	ex := errorx.New(msg).(*errorx.ErrorX)
	c := ERROR_CODE_NORMAL
	if len(code) > 0 {
		c = code[0]
	}
	return &BizError{*ex, c}
}

func PanicError(errs ...error) {
	for _, err := range errs {
		if err != nil {
			if ex, ok := err.(*errorx.ErrorX); ok {
				panic(ex)
			} else {
				panic(errorx.Stacked(err))
			}
		}
	}
}

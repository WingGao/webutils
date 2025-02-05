package werror

import "github.com/gookit/goutil/errorx"

type BizError struct {
	errorx.ErrorX
}

func NewBizError(msg string) *BizError {
	ex := errorx.New(msg).(*errorx.ErrorX)
	return &BizError{*ex}
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

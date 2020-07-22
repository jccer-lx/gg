package validation

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type GGValidationError struct {
	errMsgList []string
}

func (v *GGValidationError) Error() string {
	return strings.Join(v.errMsgList, ",")
}

func (v *GGValidationError) AppendErrMsg(errMsg string) int {
	v.errMsgList = append(v.errMsgList, errMsg)
	return len(v.errMsgList)
}

func (v *GGValidationError) GetErrMsgList() []string {
	return v.errMsgList
}

func Check(data interface{}) error {
	err := validate.Struct(data)
	ggValidationError := new(GGValidationError)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			ggValidationError.AppendErrMsg(e.Translate(trans))
		}
		return ggValidationError
	}
	return nil
}

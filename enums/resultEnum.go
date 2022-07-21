package enums

import (
	"errors"
	"fmt"
)

type errorStruct struct {
	code int
	err  error
}

func (e errorStruct) Error() string {
	return fmt.Sprintf("status %d: err %v", e.code, e.err)

}

func getParameterError() error {
	return &errorStruct{
		code: 503,
		err:  errors.New("Parameter Error!"),
	}
}

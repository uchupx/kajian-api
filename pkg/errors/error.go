package errors

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

var (
	ErrInvalidRequest = errors.New("invalid request")
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("not found")
	ErrUnauthorized   = errors.New("unauthorized")
)

type ErrorResponse struct {
	HttpCode int64  `json:"-"`
	Code     int64  `json:"code"`
	Message  string `json:"message"`
}

type ErrorMeta struct {
	Message string `json:"message"`
	Line    string `json:"line"`
	IsPanic bool   `json:"is_panic"`
}

const ErrorUnauthorizedCode = 4001
const ErrorAPIKEYNotValidCode = 4002
const ErrorDataNotFoundCode = 4003

var ErrorsMaps = map[error]ErrorResponse{}

func TracePanic() string {
	var name, file string
	var line int
	var fileLine string
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	switch {
	case name != "":
		fileLine = fmt.Sprintf("%v:%v", name, line)
	case file != "":
		fileLine = fmt.Sprintf("%v:%v", file, line)
	default:
		fileLine = fmt.Sprintf("pc:%x", pc)
	}

	return fileLine
}

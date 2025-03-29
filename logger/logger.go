package logger

import (
	"fmt"
)

type LoggerInterface interface {
	r(string, ...any)
}

type Logger struct {
}

func (l Logger) r(str string, args ...any) {
	fmt.Print("%v, %v", str, args)
}

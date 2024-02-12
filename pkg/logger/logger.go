package logger

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"
)

func LogError(err interface{}, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	switch err := err.(type) {
	case error:
		fmt.Fprintf(
			os.Stderr,
			"%s %s:%d: error: %v\n",
			time.Now().Format("2006/01/02 15:04:05"),
			file,
			line,
			fmt.Sprintf(err.Error(), args...),
		)
	case string:
		fmt.Fprintf(
			os.Stderr,
			"%s %s:%d: error: %v\n",
			time.Now().Format("2006/01/02 15:04:05"),
			file,
			line,
			fmt.Sprintf(err, args...),
		)
	}
}

func Error(err interface{}, args ...interface{}) error {
	retErrStr := ""
	switch err := err.(type) {
	case error:
		retErrStr = fmt.Sprintf(err.Error(), args...)
	case string:
		retErrStr = fmt.Sprintf(err, args...)
	}
	return errors.New(retErrStr)
}

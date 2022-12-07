package logr

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

type Valuer func() any

func hasValuer(args ...any) bool {
	for _, arg := range args {
		if _, ok := arg.(Valuer); ok {
			return true
		}
	}
	return false
}

func bindValuer(args []any) []any {
	for i, arg := range args {
		if valuer, ok := arg.(Valuer); ok {
			args[i] = valuer()
		}
	}
	return args
}

func Caller(dep int, fullpath bool) Valuer {
	return func() any {
		_, file, line, _ := runtime.Caller(dep)
		if !fullpath {
			file = filepath.Base(file)
		}
		return fmt.Sprintf("%s:%d", file, line)
	}
}

func TimeStamp() Valuer {
	return func() any {
		return time.Now().Format(time.RFC3339)
	}
}

func CallerFunc(dep int) Valuer {
	return func() any {
		pc, _, _, _ := runtime.Caller(dep)
		func_name := runtime.FuncForPC(pc).Name()
		return func_name
	}
}

package logr

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type Valuer = func() any

func hasValuer(kvs ...any) bool {
	for _, kv := range kvs {
		if _, ok := kv.(Valuer); ok {
			return true
		}
	}
	return false
}

func bindValuer(kvs ...any) []any {
	for idx, kv := range kvs {
		if valuer, ok := kv.(Valuer); ok {
			kvs[idx] = valuer()
		}
	}
	return kvs
}

func CallerFile(dep int, abspath bool) Valuer {
	return func() any {
		_, file, line, _ := runtime.Caller(dep)

		if !abspath {
			file = filepath.Base(file)
		}
		return fmt.Sprintf("%s:%d", file, line)
	}
}

func CallerFunc(dep int) Valuer {
	return func() any {
		pc, _, _, _ := runtime.Caller(dep)
		funcName := runtime.FuncForPC(pc).Name()
		return funcName
	}
}

package logr

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
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
		pc, file, line, _ := runtime.Caller(dep)
		funcName := runtime.FuncForPC(pc).Name()
		if !abspath {
			file = filepath.Base(file)

			parts := strings.Split(funcName, ".")
			funcName = parts[len(parts)-1]
		}
		return fmt.Sprintf("%s:%d#%s", file, line, funcName)
	}
}

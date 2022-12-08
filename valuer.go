package logr

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type Valuer = func(context.Context) any

func hasValuer(kvs ...any) bool {
	for _, kv := range kvs {
		if _, ok := kv.(Valuer); ok {
			return true
		}
	}
	return false
}

func bindValuer(ctx context.Context, kvs ...any) []any {
	for idx, kv := range kvs {
		if valuer, ok := kv.(Valuer); ok {
			kvs[idx] = valuer(ctx)
		}
	}
	return kvs
}

func CallerFile(dep int, abspath bool) Valuer {
	return func(ctx context.Context) any {
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

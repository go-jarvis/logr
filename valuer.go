package logr

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

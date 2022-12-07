package logr

type Vauler = func() any

func hasVauler(kvs ...any) bool {
	for _, kv := range kvs {
		if _, ok := kv.(Vauler); ok {
			return true
		}
	}
	return false
}

func bindValuer(kvs ...any) []any {
	for idx, kv := range kvs {
		if valuer, ok := kv.(Vauler); ok {
			kvs[idx] = valuer()
		}
	}
	return kvs
}

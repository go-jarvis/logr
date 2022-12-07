package logr

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(err error)
	Error(err error)

	Start()
	End()
}

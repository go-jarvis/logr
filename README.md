# logr

a golang logger library

## Usage

```go

func TestDefault(t *testing.T) {

	log := Default().SetLevel(DebugLevel)
	err := errors.New("New_ERROR")

	log = log.With(
		"kk", "vv",
		"caller", CallerFile(4, false),
		"gg",
	)

	log.Debug("number=%d", 1)
	log.Info("number=%d", 1)
	log.Warn(err)
	log.Error(err)

	ctx := WithLogger(context.Background(), log)
	subcaller(ctx)
}

func subcaller(ctx context.Context) {
	log := FromContext(ctx)

	log = log.Start() // time cost
	defer log.Stop()

	time.Sleep(532 * time.Millisecond)
	log.Info("account=%d", 100)
}
```


This code appears to be a Go (golang) package named `logr`, which defines a `Logger` interface and provides a default implementation of this interface in the form of a `levelLogger` struct. The `levelLogger` struct contains fields for a logger instance, log level, a boolean flag indicating the presence of "valuer" parameters in the logger's context, and a slice of key-value pairs that represent additional context data for the logger.

The `Logger` interface defines methods for logging messages at different levels (e.g. debug, info, warning, etc.), as well as methods for adding context data to the logger and starting/stopping a timer.

The `levelLogger` struct provides implementations for these methods. For example, the Log method takes a log level and message string as arguments, and logs the message at the specified log level. The `With` method allows additional key-value pairs to be added to the logger's context, and returns a new `levelLogger` instance with the updated context. The `Start` and `Stop` methods can be used to measure the duration of some code execution and include that duration in the log output.

Overall, this package provides a simple and flexible interface for logging in Go applications.


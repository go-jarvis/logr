package logr

import (
	"context"
	"fmt"
	"time"

	"github.com/go-jarvis/logr/slogx"
	"golang.org/x/exp/slog"
)

var _ Logger = &levelLogger{}

type levelLogger struct {
	slog  *slog.Logger
	level slog.Level

	hasValuer bool
	kvs       []any

	timer time.Time
}

func Default() Logger {
	return &levelLogger{
		slog:  slogx.Default(),
		level: slog.InfoLevel,
	}
}

type Config struct {
	level  slog.Level
	logger *slog.Logger
}

func New(c Config) Logger {
	return &levelLogger{
		level: c.level,
		slog:  c.logger,
	}
}

// Log 绑定参数，打印日志
func (log *levelLogger) Log(level slog.Level, msg string) {
	kvs := append([]any{}, log.kvs...)

	if log.hasValuer {
		kvs = bindValuer(log.slog.Context(), kvs...)
	}

	log.slog.With(kvs...).LogDepth(0, level, msg)
}

// Debug 打印 debug 日志
func (log *levelLogger) Debug(msg string, args ...any) {
	if log.Enabled(slog.DebugLevel) {
		log.Log(slog.DebugLevel, fmt.Sprintf(msg, args...))
	}
}

// Info 打印 info 日志
func (log *levelLogger) Info(msg string, args ...any) {
	if log.Enabled(slog.InfoLevel) {
		log.Log(slog.InfoLevel, fmt.Sprintf(msg, args...))
	}
}

// Warn 打印 Warn 日志
func (log *levelLogger) Warn(err error) {
	if log.Enabled(slog.WarnLevel) {
		log.Log(slog.WarnLevel, err.Error())
	}
}

// Error 打印 Error 日志
func (log *levelLogger) Error(err error) {
	if log.Enabled(slog.ErrorLevel) {
		log.Log(slog.ErrorLevel, err.Error())
	}
}

// With 添加 k=v 数据， 如果参数格式为奇数， 则在最后补充 `LACK_Unknown`
// 并返回一个新的 Logger 对象
func (log *levelLogger) With(kvs ...any) Logger {
	if len(kvs)%2 != 0 {
		kvs = append(kvs, "LACK_Unknown")
	}

	logc := log.copy()
	if !logc.hasValuer && hasValuer(kvs...) {
		logc.hasValuer = true
	}

	if log.kvs == nil {
		log.kvs = make([]any, 0)
	}

	logc.kvs = append(logc.kvs, kvs...)

	return logc
}

// Start 启动 logger 对象计时器， 计算函数使用耗时（毫秒）， 并返回一个新的 Logger 对象
//
//	log = log.Start()
//	defer log.Stop()
func (log *levelLogger) Start() Logger {
	logc := log.copy()
	logc.timer = time.Now()

	return logc
}

// Stop 停止计时， 并使用 Info 级别打印耗时（毫秒） 与 调用函数
func (log *levelLogger) Stop() {
	cost := time.Now().Sub(log.timer).Milliseconds()

	log.With(
		"cost", fmt.Sprintf("%dms", cost),
		"caller", CallerFile(5, false),
	).Info("time-cost")
}

// Enabled 比较是否符合打印日志级别
func (log *levelLogger) Enabled(level slog.Level) bool {
	return log.level <= level
}

// SetLevel 设置日志等级， 并返回一个新的 Logger 对象
func (log *levelLogger) SetLevel(level slog.Level) Logger {
	return &levelLogger{
		slog:  log.slog,
		level: level,
	}
}

// WithContext 将 context 保存到 Logger 中
func (log *levelLogger) WithContext(ctx context.Context) Logger {
	logc := log.copy()
	logc.slog = log.slog.WithContext(ctx)
	return logc
}

// Context 从 Logger 中提取 context
func (log *levelLogger) Context() context.Context {
	return log.slog.Context()
}

// copy 复制一个新的 *levelLogger 对象
func (log *levelLogger) copy() *levelLogger {
	return &levelLogger{
		slog:      log.slog,
		level:     log.level,
		hasValuer: log.hasValuer,
		kvs:       log.kvs,
	}
}

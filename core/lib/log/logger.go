package log

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	loggerSettingsJSON = `{
		"level": "debug",
		"encoding": "console",
		"outputPaths": ["stderr"],
	  	"errorOutputPaths": ["stderr"],
		"development": false,
		"disableStacktrace": true,
	  	"encoderConfig": {
	    	"messageKey": "message",
	    	"levelKey": "level",
	    	"levelEncoder": "capital",
			"timeKey": "timestamp",
			"timeEncoder": "iso8601",
			"callerKey": "caller",
			"callerEncoder": "short",
			"stacktraceKey": "stack"
	  	}
	}`

	correlationIDKey   = "correlation_id"
	requestIDKey       = "request_id"
	userAgentKey       = "user_agent"
	ipAddressKey       = "ip_address"
	originUserAgentKey = "origin_user_agent"
	pathKey            = "pathKey"
)

// Level is log level type
type Level string

// DEBUG and others are supported log levels
const (
	DEBUG  Level = "DEBUG"
	ERROR  Level = "ERROR"
	WARN   Level = "WARN"
	INFO   Level = "INFO"
	FATAL  Level = "FATAL"
	PANIC  Level = "PANIC"
	DPANIC Level = "DPANIC"
)

// Format is log format
type Format string

// CONSOLE and other are supported log formats
const (
	CONSOLE Format = "console"
	JSON    Format = "json"
)

// Log contains logger deps
type Log struct {
	Level      Level
	Time       time.Time
	LoggerName string
	Message    string
	Caller     string
	Stack      string
}

var l *zap.Logger

// nolint: gochecknoinits
func init() {
	l = New(
		Format(os.Getenv("LOG_FORMAT")),
		Level(os.Getenv("LOG_LEVEL")),
	)
}

func withGrpcContext(ctx context.Context) []zapcore.Field {
	fields := make([]zapcore.Field, 0)

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return fields
	}

	if correlationIDHeader := md.Get("x-correlation-id"); len(correlationIDHeader) == 1 {
		fields = append(fields, zap.String(correlationIDKey, correlationIDHeader[0]))
	}

	if requestIDHeader := md.Get("x-request-id"); len(requestIDHeader) == 1 {
		fields = append(fields, zap.String(requestIDKey, requestIDHeader[0]))
	}

	if UAHeader := md.Get("user-agent"); len(UAHeader) == 1 {
		fields = append(fields, zap.String(userAgentKey, UAHeader[0]))
	}

	if XFFHeader := md.Get("x-forwarded-for"); len(XFFHeader) > 0 {
		fields = append(fields, zap.String(ipAddressKey, XFFHeader[0]))
	}

	if XFUAHeader := md.Get("x-forwarded-user-agent"); len(XFUAHeader) > 0 {
		fields = append(fields, zap.String(originUserAgentKey, XFUAHeader[0]))
	}

	method, ok := grpc.Method(ctx)

	if ok {
		// Method looks something like - protos.Api/Rpc
		splitMethod := strings.Split(method, ".")

		// Get the last value - Api/Rpc
		method = splitMethod[len(splitMethod)-1]
		// Do a char replacement to . to properly index - Api.Rpc
		method = strings.ReplaceAll(method, "/", ".")

		fields = append(fields, zap.String(pathKey, method))
	}

	return fields
}

func withContext(ctx context.Context) []zapcore.Field {
	fields := make([]zapcore.Field, 0)
	fields = append(fields, withGrpcContext(ctx)...)

	return fields
}

func withStackTrace(fields []zapcore.Field) []zapcore.Field {
	return append(fields, zap.Stack("stacktrace"))
}

// Field is a log message key value pair
type Field struct {
	Key   string
	Value interface{}
}

// Options contains logger options
type Options struct {
	Name     *string         `json:"name,omitempty"`
	Message  string          `json:"message"`
	LogLevel Level           `json:"log_level"`
	Tag      *string         `json:"tag,omitempty"`
	Hostname *string         `json:"hostname,omitempty"`
	LogTime  *time.Time      `json:"log_time,omitempty"`
	Err      error           `json:"error"`
	Fields   []Field         `json:"fields"`
	Context  context.Context `json:"context"`
}

// Option is a logger optional
type Option func(*Options)

func newOptionFields(ctx context.Context, msg string, logLvl Level, errOpts ...Option) []zapcore.Field {
	fields := withContext(ctx)
	opts := &Options{
		Message:  msg,
		Context:  ctx,
		LogLevel: logLvl,
	}

	for _, option := range errOpts {
		option(opts)
	}

	if opts.Err != nil {
		notifySentry(opts.Err)
		fields = append(fields, zap.Error(opts.Err))
	}

	for _, field := range opts.Fields {
		fields = append(fields, zap.Reflect(field.Key, field.Value))
	}

	return fields
}

// WithValue appends a value to log message
func WithValue(key string, value interface{}) Option {
	return func(opts *Options) {
		opts.Fields = append(opts.Fields, Field{
			Key:   key,
			Value: value,
		})
	}
}

// WithError appends an error to the log message
func WithError(err error) Option {
	return func(opts *Options) {
		opts.Err = err
	}
}

// New inits a new logger
func New(env Format, logLevel Level) *zap.Logger {
	var cfg zap.Config
	if err := json.Unmarshal([]byte(loggerSettingsJSON), &cfg); err != nil {
		panic(err)
	}

	if env == JSON {
		cfg.Encoding = "json"
	}

	switch logLevel {
	case DEBUG:
		cfg.Level.SetLevel(zap.DebugLevel)
	case ERROR:
		cfg.Level.SetLevel(zap.ErrorLevel)
	case WARN:
		cfg.Level.SetLevel(zap.WarnLevel)
	case INFO:
		cfg.Level.SetLevel(zap.InfoLevel)
	default:
		cfg.Level.SetLevel(zap.InfoLevel)
	}

	logger, err := cfg.Build()

	if err != nil {
		panic(err)
	}

	l = logger.WithOptions(zap.AddCallerSkip(1))

	return l
}

// Debug log
func Debug(ctx context.Context, msg string, opts ...Option) {
	l.Debug(msg, newOptionFields(ctx, msg, DEBUG, opts...)...)
}

// Info log
func Info(ctx context.Context, msg string, opts ...Option) {
	l.Info(msg, newOptionFields(ctx, msg, INFO, opts...)...)
}

// Warn log
func Warn(ctx context.Context, msg string, opts ...Option) {
	l.Warn(msg, newOptionFields(ctx, msg, WARN, opts...)...)
}

// Error log
func Error(ctx context.Context, msg string, opts ...Option) {
	l.Error(msg, withStackTrace(newOptionFields(ctx, msg, ERROR, opts...))...)
}

// Fatal log
func Fatal(ctx context.Context, msg string, opts ...Option) {
	l.Fatal(msg, withStackTrace(newOptionFields(ctx, msg, FATAL, opts...))...)
}

// Panic log
func Panic(ctx context.Context, msg string, opts ...Option) {
	l.Panic(msg, withStackTrace(newOptionFields(ctx, msg, PANIC, opts...))...)
}

// DPanic log
func DPanic(ctx context.Context, msg string, opts ...Option) {
	l.DPanic(msg, withStackTrace(newOptionFields(ctx, msg, DPANIC, opts...))...)
}

// Sync ?
func Sync() {
	_ = l.Sync()
}

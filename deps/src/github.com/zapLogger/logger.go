package zapLogger

import (
	"errors"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// error zapLogger
var sugaredLogger *zap.SugaredLogger

var timeFormat = "2006-01-02 15:04:05.999999999 -0700 MST"

// 是否是trace级别的输出
var isTraceLevel = false

var levelMap = map[string]zapcore.Level{
	"trace":  zapcore.DebugLevel,
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

type LogConfig struct {
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
	Compress   bool
	Level      string
}

var Logc LogConfig

var logLevel = [6]string{"trace", "debug", "info", "warn", "error", "fatal"}

func GetLoggerLevel(index int) (level string, err error) {
	if index+1 >= len(logLevel) {
		return "", errors.New("log index error")
	}

	level = logLevel[index+1]
	return level, nil
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func init() {
	logc := LogConfig{
		Filename:   "/var/log/zap.log", // 日志文件路径
		MaxSize:    1,                  // megabytes(兆字节)
		MaxBackups: 7,                  // 最多保留3个备份
		MaxAge:     3,                  // days
		Compress:   true,               // 是否压缩 disabled by default
		LocalTime:  true,               // 用于格式化备份文件中的时间戳的时间是计算机的本地时间。 默认是使用UTC
	}
	_ = InitLogger("info", logc)
}

func InitLogger(level string, logc LogConfig) error {
	if logc.Filename == "" {
		return errors.New("invalid LogConfig")
	}
	Logc = logc
	log := lumberjack.Logger{
		Filename:   logc.Filename,   // 日志文件路径
		MaxSize:    logc.MaxSize,    // megabytes(兆字节)
		MaxBackups: logc.MaxBackups, // 最多保留3个备份
		MaxAge:     logc.MaxAge,     // days
		Compress:   logc.Compress,   // 是否压缩 disabled by default
		LocalTime:  logc.LocalTime,  // 用于格式化备份文件中的时间戳的时间是计算机的本地时间。 默认是使用UTC
	}

	lv := getLoggerLevel(level)
	if "trace" == level {
		isTraceLevel = true
	}

	syncWriter := zapcore.AddSync(&log)
	encoder := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",                          // json时时间键
		LevelKey:       "L",                          // json时日志等级键
		NameKey:        "N",                          // json时日志记录器键
		CallerKey:      "C",                          // json时日志文件信息键
		MessageKey:     "M",                          // json时日志消息键
		StacktraceKey:  "S",                          // json时堆栈键
		LineEnding:     zapcore.DefaultLineEnding,    // 友好日志换行符
		EncodeLevel:    zapcore.CapitalLevelEncoder,  // 友好日志等级名大小写（info INFO）
		EncodeTime:     TimeEncoder,                  // 友好日志时日期格式化
		EncodeDuration: zapcore.NanosDurationEncoder, // 时间序列化
		EncodeCaller:   zapcore.ShortCallerEncoder,   // 日志文件信息（包/文件.go:行号）
	}
	core := zapcore.NewTee(
		// 有好的格式、输出控制台、动态等级
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), os.Stdout, lv),
		// 有好的格式、输出文件、处定义等级规则
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncWriter, lv),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugaredLogger = logger.Sugar()
	return nil
}

// TimeEncoder  格式化时间
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(timeFormat))
}

// Trace trace输出
func Trace(args ...interface{}) bool {
	if isTraceLevel {
		sugaredLogger.Debug(args...)
	}
	return isTraceLevel
}

// Tracef Tracef输出
func Tracef(template string, args ...interface{}) bool {
	if isTraceLevel {
		sugaredLogger.Debugf(template, args...)
	}
	return isTraceLevel
}

// Debug  debug输出
func Debug(args ...interface{}) {
	sugaredLogger.Debug(args...)
}

// Debugf Debug输出
func Debugf(template string, args ...interface{}) {
	sugaredLogger.Debugf(template, args...)
}

// Info Info输出
func Info(args ...interface{}) {
	sugaredLogger.Info(args...)
}

// Infof Trace
func Infof(template string, args ...interface{}) {
	sugaredLogger.Infof(template, args...)
}

// Warn warm输出
func Warn(args ...interface{}) {
	sugaredLogger.Warn(args...)
}

// Warnf Warn输出
func Warnf(template string, args ...interface{}) {
	sugaredLogger.Warnf(template, args...)
}

// Error error输出
func Error(args ...interface{}) {
	sugaredLogger.Error(args...)
}

// Errorf error输出
func Errorf(template string, args ...interface{}) {
	sugaredLogger.Errorf(template, args...)
}

// DPanic DPanic输出
func DPanic(args ...interface{}) {
	sugaredLogger.DPanic(args...)
}

// DPanicf DPanic输出
func DPanicf(template string, args ...interface{}) {
	sugaredLogger.DPanicf(template, args...)
}

// Panic Panic输出
func Panic(args ...interface{}) {
	sugaredLogger.Panic(args...)
}

// Panicf Panicf输出
func Panicf(template string, args ...interface{}) {
	sugaredLogger.Panicf(template, args...)
}

// Fatal Fatal输出
func Fatal(args ...interface{}) {
	sugaredLogger.Fatal(args...)
}

// Fatalf Fatalf输出
func Fatalf(template string, args ...interface{}) {
	sugaredLogger.Fatalf(template, args...)
}

// Logger 给IRIS装载自定义的Log库
type Logger struct {
}

// Debug Debug
func (l *Logger) Debug(args ...interface{}) {
	sugaredLogger.Debug(args...)
}

// Info Info
func (l *Logger) Info(args ...interface{}) {
	sugaredLogger.Info(args...)
}

// Warn Warn
func (l *Logger) Warn(args ...interface{}) {
	sugaredLogger.Warn(args...)
}

// Error Error
func (l *Logger) Error(args ...interface{}) {
	sugaredLogger.Error(args...)
}

// Println Println
func (l *Logger) Println(args ...interface{}) {
	sugaredLogger.Debug(args...)
}

// Print Print
func (l *Logger) Print(args ...interface{}) {
	sugaredLogger.Debug(args...)
}

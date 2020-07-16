package zapLogger

// import (
// 	"testing"

// 	"git.vnnox.net/controller/ctl/config/systemconfig"
// 	"gopkg.in/natefinch/lumberjack.v2"
// )

// func BenchmarkInitLogger(b *testing.B) {
// 	b.ResetTimer()
// 	b.ReportAllocs()
// 	hook := lumberjack.Logger{
// 		Filename:   "./zap.log",                      // 日志文件路径
// 		MaxSize:    systemconfig.Data.Log.MaxSize,    // megabytes(兆字节)
// 		MaxBackups: systemconfig.Data.Log.MaxBackups, // 最多保留3个备份
// 		MaxAge:     systemconfig.Data.Log.MaxAge,     // days
// 		Compress:   systemconfig.Data.Log.Compress,   // 是否压缩 disabled by default
// 		LocalTime:  systemconfig.Data.Log.LocalTime,  // 用于格式化备份文件中的时间戳的时间是计算机的本地时间。 默认是使用UTC
// 	}
// 	_ = InitLogger("trace", &hook)
// 	Debug(">>>>>> End InitLogger 1")

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			Trace("Trace i:", i)
// 			Debug("Debug i:", i)
// 			Warn("Warn i:", i)
// 			Error("Error:", i)
// 		}
// 	}()

// 	hook1 := lumberjack.Logger{
// 		Filename:   "./zap.log",                      // 日志文件路径
// 		MaxSize:    systemconfig.Data.Log.MaxSize,    // megabytes(兆字节)
// 		MaxBackups: systemconfig.Data.Log.MaxBackups, // 最多保留3个备份
// 		MaxAge:     systemconfig.Data.Log.MaxAge,     // days
// 		Compress:   systemconfig.Data.Log.Compress,   // 是否压缩 disabled by default
// 		LocalTime:  systemconfig.Data.Log.LocalTime,  // 用于格式化备份文件中的时间戳的时间是计算机的本地时间。 默认是使用UTC
// 	}
// 	_ = InitLogger("trace", &hook1)
// 	Debug(">>>>>> End InitLogger 2")
// }

// func TestInitLogger(t *testing.T) {
// 	var i int
// 	for i < 1000 {
// 		hook := lumberjack.Logger{
// 			Filename:   "./zap.log",                      // 日志文件路径
// 			MaxSize:    systemconfig.Data.Log.MaxSize,    // megabytes(兆字节)
// 			MaxBackups: systemconfig.Data.Log.MaxBackups, // 最多保留3个备份
// 			MaxAge:     systemconfig.Data.Log.MaxAge,     // days
// 			Compress:   systemconfig.Data.Log.Compress,   // 是否压缩 disabled by default
// 			LocalTime:  systemconfig.Data.Log.LocalTime,  // 用于格式化备份文件中的时间戳的时间是计算机的本地时间。 默认是使用UTC
// 		}
// 		_ = InitLogger("trace", &hook)
// 		i++
// 		Debug(">>>>>> End InitLogger 1")
// 	}
// }

// func BenchmarkDebug(b *testing.B) {
// 	hook := lumberjack.Logger{
// 		Filename:   "./zap.log",                      // 日志文件路径
// 		MaxSize:    systemconfig.Data.Log.MaxSize,    // megabytes(兆字节)
// 		MaxBackups: systemconfig.Data.Log.MaxBackups, // 最多保留3个备份
// 		MaxAge:     systemconfig.Data.Log.MaxAge,     // days
// 		Compress:   systemconfig.Data.Log.Compress,   // 是否压缩 disabled by default
// 		LocalTime:  systemconfig.Data.Log.LocalTime,  // 用于格式化备份文件中的时间戳的时间是计算机的本地时间。 默认是使用UTC
// 	}
// 	_ = InitLogger("trace", &hook)
// 	b.ResetTimer()
// 	Debug("next package")
// 	b.ReportAllocs()
// }

// func BenchmarkInfo(b *testing.B) {
// 	hook := lumberjack.Logger{
// 		Filename:   "./zap.log",                      // 日志文件路径
// 		MaxSize:    systemconfig.Data.Log.MaxSize,    // megabytes(兆字节)
// 		MaxBackups: systemconfig.Data.Log.MaxBackups, // 最多保留3个备份
// 		MaxAge:     systemconfig.Data.Log.MaxAge,     // days
// 		Compress:   systemconfig.Data.Log.Compress,   // 是否压缩 disabled by default
// 		LocalTime:  systemconfig.Data.Log.LocalTime,  // 用于格式化备份文件中的时间戳的时间是计算机的本地时间。 默认是使用UTC
// 	}
// 	_ = InitLogger("trace", &hook)
// 	b.N = 100
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		Debug("next package asdfsdfasdsdfsdfasdfasfasfasf sdfsdafasdfasf sdfdsad asd safasdf")
// 	}
// }

// func BenchmarkInfo1(b *testing.B) {
// 	hook := lumberjack.Logger{
// 		Filename:   "./zap.log",                // 日志文件路径
// 		MaxSize:    config.Data.Log.MaxSize,    // megabytes(兆字节)
// 		MaxBackups: config.Data.Log.MaxBackups, // 最多保留3个备份
// 		MaxAge:     config.Data.Log.MaxAge,     // days
// 		Compress:   config.Data.Log.Compress,   // 是否压缩 disabled by default
// 		LocalTime:  config.Data.Log.LocalTime,  // 用于格式化备份文件中的时间戳的时间是计算机的本地时间。 默认是使用UTC
// 	}
// 	_ = InitLogger("trace", &hook)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		Debug("next package asdfsdfasdsdfsdfasdfasfasfasf sdfsdafasdfasf sdfdsad asd safasdf")
// 	}
// }

// func BenchmarkInfo2(b *testing.B) {
// 	hook := lumberjack.Logger{
// 		Filename:   "./zap.log",                // 日志文件路径
// 		MaxSize:    config.Data.Log.MaxSize,    // megabytes(兆字节)
// 		MaxBackups: config.Data.Log.MaxBackups, // 最多保留3个备份
// 		MaxAge:     config.Data.Log.MaxAge,     // days
// 		Compress:   config.Data.Log.Compress,   // 是否压缩 disabled by default
// 		LocalTime:  config.Data.Log.LocalTime,  // 用于格式化备份文件中的时间戳的时间是计算机的本地时间。 默认是使用UTC
// 	}
// 	_ = InitLogger("trace", &hook)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		Debug("next package asdfsdfasdsdfsdfasdfasfasfasf sdfsdafasdfasf sdfdsad asd safasdf")
// 	}
// }

// func BenchmarkTimeEncoder(b *testing.B) {
// 	hook := lumberjack.Logger{
// 		Filename:   "./zap.log",                      // 日志文件路径
// 		MaxSize:    systemconfig.Data.Log.MaxSize,    // megabytes(兆字节)
// 		MaxBackups: systemconfig.Data.Log.MaxBackups, // 最多保留3个备份
// 		MaxAge:     systemconfig.Data.Log.MaxAge,     // days
// 		Compress:   systemconfig.Data.Log.Compress,   // 是否压缩 disabled by default
// 		LocalTime:  systemconfig.Data.Log.LocalTime,  // 用于格式化备份文件中的时间戳的时间是计算机的本地时间。 默认是使用UTC
// 	}
// 	_ = InitLogger("info", &hook)

// 	b.N = 100
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		tr1()
// 	}
// }

// func tr1() {
// 	if Trace("sdfsdfsdfsdfsdfsdfsdsfsdfsdfsdfsdfsdf") {
// 		defer Trace("sdfsdfsdfsdfdsfsdfsdf")
// 	}
// }

// func tr2() {
// 	if Trace("sdfsdfsdfsdfsdfsdfsdsfsdfsdfsdfsdfsdf") {
// 		defer Trace("sdfsdfsdfsdfdsfsdfsdf")
// 	}
// }

// func tr3() {
// 	if Trace("sdfsdfsdfsdfsdfsdfsdsfsdfsdfsdfsdfsdf") {
// 		defer Trace("sdfsdfsdfsdfdsfsdfsdf")
// 	}
// }

// func tr4() {
// 	if Trace("sdfsdfsdfsdfsdfsdfsdsfsdfsdfsdfsdfsdf") {
// 		defer Trace("sdfsdfsdfsdfdsfsdfsdf")
// 	}
// }

// func tr5() {
// 	if Trace("sdfsdfsdfsdfsdfsdfsdsfsdfsdfsdfsdfsdf") {
// 		defer Trace("sdfsdfsdfsdfdsfsdfsdf")
// 	}
// }

// func tr6() {
// 	if Trace("sdfsdfsdfsdfsdfsdfsdsfsdfsdfsdfsdfsdf") {
// 		defer Trace("sdfsdfsdfsdfdsfsdfsdf")
// 	}
// }

// func BenchmarkSprint(b *testing.B) {
// }

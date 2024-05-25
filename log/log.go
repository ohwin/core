package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLogger() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder, // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	// 设置日志级别
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel), // 日志级别
		Development:      true,                                 // 开发模式，堆栈跟踪
		Encoding:         "console",                            // 输出格式 console 或 yaml
		EncoderConfig:    encoderConfig,                        // 编码器配置
		OutputPaths:      []string{"stdout"},                   // 输出到指定文档 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}
	zap.AddStacktrace(zapcore.FatalLevel)
	const omitFrames = 2

	log, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return
	}
	logger = log
}

func Error(msg string, param ...any) {
	logger.Error(fmt.Sprintf(msg, param...))
}

func Warn(msg string, param ...any) {
	logger.Warn(fmt.Sprintf(msg, param...))
}

func Debug(msg string, param ...any) {
	logger.Debug(fmt.Sprintf(msg, param...))
}

func Info(msg string, param ...any) {
	logger.Info(fmt.Sprintf(msg, param...))
}

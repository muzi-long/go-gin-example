package logger

import (
	"os"
	"path"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(c *Config) (*zap.Logger, error) {
	hook := lumberjack.Logger{
		Filename:   path.Join(c.Path, c.FileName), // 日志文件路径，默认 os.TempDir()
		MaxSize:    c.MaxSize,                     // 每个日志文件保存10M，默认 100M
		MaxBackups: c.MaxBackups,                  // 保留30个备份，默认不限
		MaxAge:     c.MaxAge,                      // 保留7天，默认不限
		Compress:   c.Compress,                    // 是否压缩，默认不压缩
	}
	// 设置日志级别
	var level zapcore.Level
	switch c.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.DebugLevel
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		level,
	)
	gLogger := zap.New(core, zap.AddCaller(), zap.Development())
	return gLogger, nil
}

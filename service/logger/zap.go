package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type zapLog struct {
	logger *zap.Logger
}

func NewZapLog() *zapLog {
	encoder := getEncoder()
	infoDebugWarnWriter := getInfoDebugWarnLogWriter()
	infoDebugWarnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel || lvl == zapcore.DebugLevel || lvl == zapcore.WarnLevel
	})
	infoDebugWarnCore := zapcore.NewCore(encoder, infoDebugWarnWriter, infoDebugWarnLevel)
	errorWriter := getErrorLogWriter()
	errorLevel := zap.ErrorLevel
	errorCore := zapcore.NewCore(encoder, errorWriter, errorLevel)
	core := zapcore.NewTee(infoDebugWarnCore, errorCore)
	return &zapLog{
		logger: zap.New(core, zap.AddCaller()),
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// JsonEncoder 输出日志为json格式
//func getJsonEncoder() zapcore.Encoder {
//	encoderConfig := zap.NewProductionEncoderConfig()
//	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
//	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
//	return zapcore.NewJSONEncoder(encoderConfig)
//}

func getInfoDebugWarnLogWriter() zapcore.WriteSyncer {
	fileName := viper.GetString("log.logFile")
	if fileName == "" {
		panic("getInfoDebugWarnLogWriter error: no log.logFile in config file")
	}
	maxAge := viper.GetInt("log.maxAge")
	if maxAge == 0 {
		maxAge = 30
	}
	rotateTime := viper.GetInt("log.rotateTime")
	if rotateTime == 0 {
		rotateTime = 1
	}
	hook, err := rotatelogs.New(
		fileName+"_%Y%m%d%H.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(time.Hour*24*time.Duration(maxAge)),
		rotatelogs.WithRotationTime(time.Hour*time.Duration(rotateTime)),
	)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(hook)
}

func getErrorLogWriter() zapcore.WriteSyncer {
	fileName := viper.GetString("log.errorLog")
	if fileName == "" {
		panic("getErrorLogWriter error: no log.errorLog in config file")
	}
	file, _ := os.OpenFile(fileName+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
	return zapcore.AddSync(file)
}

func (z *zapLog) Info(msg string) {
	z.logger.Info(msg)
}

func (z *zapLog) Error(msg string) {
	z.logger.Error(msg)
}

func (z *zapLog) Debug(msg string) {
	z.logger.Debug(msg)
}

func (z *zapLog) Warn(msg string) {
	z.logger.Warn(msg)
}

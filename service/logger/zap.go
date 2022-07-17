package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"pandora/service/config"
	"time"
)

type zapLog struct {
	logger *zap.Logger
	conf   config.Config
}

func NewZapLog(conf config.Config) *zapLog {
	zl := &zapLog{
		conf: conf,
	}
	encoder := zl.getEncoder()
	infoDebugWarnWriter := zl.getInfoDebugWarnLogWriter()
	infoDebugWarnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel || lvl == zapcore.DebugLevel || lvl == zapcore.WarnLevel
	})
	infoDebugWarnCore := zapcore.NewCore(encoder, infoDebugWarnWriter, infoDebugWarnLevel)
	errorWriter := zl.getErrorLogWriter()
	errorLevel := zap.ErrorLevel
	errorCore := zapcore.NewCore(encoder, errorWriter, errorLevel)
	core := zapcore.NewTee(infoDebugWarnCore, errorCore)
	zl.logger = zap.New(core, zap.AddCaller())
	return zl
}

func (zl *zapLog) getEncoder() zapcore.Encoder {
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

func (zl *zapLog) getInfoDebugWarnLogWriter() zapcore.WriteSyncer {
	fileName := zl.conf.GetString("log.logFile")
	if fileName == "" {
		panic("getInfoDebugWarnLogWriter error: no log.logFile in config file")
	}
	maxAge := zl.conf.GetInt("log.maxAge")
	if maxAge == 0 {
		maxAge = 30
	}
	rotateTime := zl.conf.GetInt("log.rotateTime")
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

func (zl *zapLog) getErrorLogWriter() zapcore.WriteSyncer {
	fileName := zl.conf.GetString("log.errorLog")
	if fileName == "" {
		panic("getErrorLogWriter error: no log.errorLog in config file")
	}
	file, _ := os.OpenFile(fileName+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
	return zapcore.AddSync(file)
}

func (zl *zapLog) Info(msg string) {
	zl.logger.Info(msg)
}

func (zl *zapLog) Error(msg string) {
	zl.logger.Error(msg)
}

func (zl *zapLog) Debug(msg string) {
	zl.logger.Debug(msg)
}

func (zl *zapLog) Warn(msg string) {
	zl.logger.Warn(msg)
}

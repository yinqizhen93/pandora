package service

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Logger *zap.Logger

//func main() {
//	InitLogger()
//	defer Logger.Sync()
//	simpleHttpGet("www.sogo.com")
//	simpleHttpGet("http://www.sogo.com")
//}

func InitLogger() *zap.Logger {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("日志初始化失败")
			panic(fmt.Sprintf("日志初始化失败: %s", err))
		} else {
			fmt.Println("log初始化成功...")
		}
	}()
	infoWriteSyncer := getInfoLogWriter()
	encoder := getEncoder()
	jsonEncoder := getJsonEncoder()
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	infoCore := zapcore.NewCore(jsonEncoder, infoWriteSyncer, infoLevel)
	errWriteSyncer := getErrorLogWriter()
	errorCore := zapcore.NewCore(encoder, errWriteSyncer, zapcore.ErrorLevel)
	core := zapcore.NewTee(infoCore, errorCore)
	Logger = zap.New(core, zap.AddCaller())
	//fmt.Println("log初始化成功...")
	return Logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getJsonEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getInfoLogWriter() zapcore.WriteSyncer {
	fileName := viper.GetString("log.logFile")
	maxAge := viper.GetInt("log.maxAge")
	rotateTime := viper.GetInt("log.rotateTime")
	// todo 检查配置变量是否存在 fileName == ""
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
	file, _ := os.OpenFile(fileName+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
	writeSyncer := zapcore.AddSync(file)
	return writeSyncer
}

//func SimpleHttpGet(url string) {
//	Logger.Debug("Trying to hit GET request for %s")
//	resp, err := http.Get(url)
//	if err != nil {
//		//sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
//		Logger.Error("Error fetching URL",
//			zap.String("name", "yinqizhen"),
//			zap.Int("OrgId", 233))
//	} else {
//		//sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
//		Logger.Info("Success! statusCode",
//			zap.String("name", "yinqizhen"),
//			zap.Int("OrgId", 233))
//		resp.Body.Close()
//	}
//}

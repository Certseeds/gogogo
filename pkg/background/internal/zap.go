// SPDX-License-Identifier: AGPL-3.0-or-later
package internal

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Logger *zap.SugaredLogger

func init() {
	l, _ := rotatelogs.New(
		"./background.log"+".%Y%m%d",
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 最长保存30天
		rotatelogs.WithRotationTime(time.Hour*24), // 24小时切割一次
	)
	zapcore.AddSync(l)
	Logger = getLogger()
	Logger.Info("init finish")
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	file, _ := os.Create(fmt.Sprintf("./%s", filename))
	return zapcore.AddSync(file)
}
func getLogger() *zap.SugaredLogger {
	encoder := getEncoder()

	logF := getLogWriter("./background.log")
	c1 := zapcore.NewCore(encoder, zapcore.AddSync(logF), zapcore.DebugLevel)
	// test.err.log记录ERROR级别的日志
	errF := getLogWriter("./background.err.log")
	c2 := zapcore.NewCore(encoder, zapcore.AddSync(errF), zap.ErrorLevel)
	// 使用NewTee将c1和c2合并到core
	core := zapcore.NewTee(c1, c2)

	logger := zap.New(core, zap.AddCaller())
	sugar := logger.Sugar()
	return sugar
}

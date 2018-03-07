// 本示例主要用于展示如何为一个logger实例配置一个或多个writer。

package main

import (
	"fmt"
	"time"

	"github.com/sevenNt/wzap"
)

// defaultLogger 使用vega/log包默认logger输出日志。
func defaultLogger() {
	// vega/log包中默认logger实例仅包含consoleWriter，
	// 开发者可以使用wzap.Info等方法使用该logger实例输出日志。
	wzap.Debug("default log debug")
	wzap.Info("default log info")
	wzap.Warn("default log warn")
	wzap.Error("default log error")

	// 构造使用fileWriter的logger实例，并覆盖vega/log包的默认logger后，
	// 开发者可以使用wzap.Info等方法，来使用覆盖后的logger输出日志。
	logger := wzap.New(
		wzap.WithLevel(wzap.Info),
		wzap.WithPath("./defaultLogger.log"),
		wzap.WithFields(wzap.Int("hahaah", 10), wzap.String("dadsada", "fafasfa")),
	)
	// 使用SetDefaultLogger方法将指定logger实例注入到vega/log包中。
	// 使用wzap.Debug等方法会调用注入的logger实例输出日志。
	wzap.SetDefaultLogger(logger)
	wzap.Debug("debug")
	wzap.Info("info", "name", 123)
	wzap.Warn("warn")
	wzap.Error("error")
}

// fileWriterLogger 仅使用fileWriter写日志。
func fileWriterLogger() {
	logger := wzap.New(
		wzap.WithLevelCombo("Warn | Error | Panic | Fatal"), // 只有级别为Warn或Error、Panic、Fatal日志才会被写入。
		wzap.WithPath("./fileWriterLogger.log"),
	)
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}

// multiWriterLogger 使用多个writer(两个fileWriter、一个consoleWriter)同时写日志。
func multiWriterLogger() {
	wzap.SetDefaultFields(wzap.Int("appid", 100010), wzap.String("appname", "test-go"))
	logger := wzap.New(
		wzap.WithOutput(
			wzap.WithLevel(wzap.Info),
			wzap.WithPath("./multiWriterLogger1.log"),
		),
		wzap.WithOutput(
			wzap.WithLevelString("Warn | Error | Panic | Fatal"),
			wzap.WithPath("./multiWriterLogger2.log"),
		),
		wzap.WithOutput(
			wzap.WithLevelMask(wzap.DebugLevel),
			wzap.WithColorful(true),
		),
		wzap.WithLevelMask(wzap.InfoLevel|wzap.WarnLevel|wzap.FatalLevel|wzap.ErrorLevel),
		wzap.WithPath("./multiWriterLogger.log"),
	)
	logger.Debug("debug", "aaa", 123, "bbb", 1234)
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}

func main() {
	fmt.Println(wzap.DebugLevel, wzap.InfoLevel, wzap.WarnLevel, wzap.ErrorLevel, wzap.FatalLevel)
	multiWriterLogger()
	fileWriterLogger()
	defaultLogger()

	time.Sleep(time.Second)
}

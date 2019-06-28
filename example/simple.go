// 本示例主要用于展示如何为一个logger实例配置一个或多个writer。

package main

import (
	"time"

	"github.com/sevennt/wzap"
)

func main() {
	// 使用SetDefaultFields方法新增全局默认filed，所有logger实例记录日志时都将包含该field。
	wzap.SetDefaultFields(
		wzap.String("iid", "test12313131231451"),
	)
	wzap.SetDefaultDir("./log/")
	logger := wzap.New(
		wzap.WithLevel(wzap.Info),
		wzap.WithPath("simple.log"),
		// 使用WithFields方法对指定logger实例新增默认filed，该logger实例记录日志时都将包含该field。
		wzap.WithFields(wzap.Int("key1", 10), wzap.String("dadsada", "fafasfa")),
		wzap.WithOutput(
			wzap.WithLevelMask(wzap.DebugLevel),
			wzap.WithColorful(true),
			wzap.WithFields(wzap.Int("key2", 10), wzap.String("dadsada", "fafasfa")),
		),
	)
	wzap.SetDefaultLogger(logger)
	wzap.Debug("debug")
	wzap.Info("info")
	wzap.Warn("warn")
	wzap.Error("error")

	time.Sleep(time.Second)
}

package wzap_test

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"testing"

	"github.com/sevenNt/wzap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// go test github.com/sevenNt/wzap -v -run=TestFileWrite$
func TestFileWrite(t *testing.T) {
	temp, err := ioutil.TempFile("/tmp", "async")
	require.NoError(t, err, "Failed to create temp file.")

	defer os.Remove(temp.Name())
	logger := wzap.New(
		wzap.WithPath(temp.Name()), // 日志写入文件为/tmp/async.log
		wzap.WithLevel(wzap.Error), // 日志写入级别为Error，此时只有级别大于等于Error级别日志才会被写入
	)
	assert.NoError(t, err)

	// 写入Info级别日志（由于本级别小小于Error级别，此日志不会被写入）
	logger.Info("some info message",
		"scenario", "classroom",
		"Hanmeimei", "How are you?",
		"Lilei", "I'm fine, thank you.",
	)
	// 写入Error级别日志（此级别大于等于Error级别，会被写入）
	logger.Error("some error message",
		"scenario", "classroom",
		"Hanmeimei", "How are you?",
		"Lilei", "I'm fine, thank you.",
	)

	countMsg := "sampling"
	expectCount := 200
	for i := 0; i < expectCount; i++ {
		logger.Error(countMsg)
	}

	byteContents, err := ioutil.ReadAll(temp)
	require.NoError(t, err, "Couldn't read log contents from temp file.")
	logs := string(byteContents)

	assert.Contains(t, logs, `"scenario":"classroom","Hanmeimei":"How are you?","Lilei":"I'm fine, thank you."`, "Unexpected log output.")
	assert.Equal(t, expectCount, strings.Count(logs, countMsg))
}

func BenchmarkStructedLogError(b *testing.B) {
	temp, err := ioutil.TempFile("/tmp", "async")
	defer os.Remove(temp.Name())
	require.NoError(b, err, "Failed to create temp file.")

	logger := wzap.New(
		wzap.WithPath(temp.Name()), // 日志写入文件为/tmp/async.log
		wzap.WithLevel(wzap.Error), // 日志写入级别为Error，此时只有级别大于等于Error级别日志才会被写入
	)
	for i := 0; i < b.N; i++ {
		logger.Error("random int bench",
			"random", rand.Intn(100),
		)
	}
}

func BenchmarkStructedLogInfo(b *testing.B) {
	temp, err := ioutil.TempFile("/tmp", "async")
	defer os.Remove(temp.Name())
	require.NoError(b, err, "Failed to create temp file.")

	logger := wzap.New(
		wzap.WithPath(temp.Name()), // 日志写入文件为/tmp/async.log。
		wzap.WithLevel(wzap.Info),  // 日志写入级别为Error，此时只有级别大于等于Error级别日志才会被写入。
	)
	for i := 0; i < b.N; i++ {
		logger.Info("random int bench",
			"random", rand.Intn(100),
		)
	}
}

func BenchmarkConsoleLogInfo(b *testing.B) {
	logger := wzap.New(
		wzap.WithLevel(wzap.Debug),
		wzap.WithColorful(true),
		wzap.WithAsync(false),
		wzap.WithPrefix("APP]>"),
	)
	for i := 0; i < b.N; i++ {
		logger.Info("random int bench",
			"random", rand.Intn(100),
		)
	}
}

func BenchmarkPrintfLogError(b *testing.B) {
	temp, err := ioutil.TempFile("/tmp", "async")
	defer os.Remove(temp.Name())
	require.NoError(b, err, "Failed to create temp file.")

	logger := wzap.New(
		wzap.WithPath(temp.Name()), // 日志写入文件为/tmp/async.log
		wzap.WithLevel(wzap.Error), // 日志写入级别为Error，此时只有级别大于等于Error级别日志才会被写入
	)
	for i := 0; i < b.N; i++ {
		logger.Errorf("random_int:%d", rand.Intn(100))
	}
}

func BenchmarkPrintfLogInfo(b *testing.B) {
	temp, err := ioutil.TempFile("/tmp", "async")
	defer os.Remove(temp.Name())
	require.NoError(b, err, "Failed to create temp file.")

	logger := wzap.New(
		wzap.WithPath(temp.Name()), // 日志写入文件为/tmp/async.log
		wzap.WithLevel(wzap.Info),  // 日志写入级别为Error，此时只有级别大于等于Error级别日志才会被写入
	)
	for i := 0; i < b.N; i++ {
		logger.Infof("random_int:%d", rand.Intn(100))
	}
}

func BenchmarkZapLogError(b *testing.B) {
	temp, err := ioutil.TempFile("/tmp", "async")
	defer os.Remove(temp.Name())
	require.NoError(b, err, "Failed to create temp file.")

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{temp.Name()}
	myLog, _ := cfg.Build()
	for i := 0; i < b.N; i++ {
		myLog.Sugar().Errorw(
			"random int bench",
			"random", rand.Intn(100),
		)
	}
}

func BenchmarkZapLogInfo(b *testing.B) {
	temp, err := ioutil.TempFile("/tmp", "async")
	defer os.Remove(temp.Name())
	require.NoError(b, err, "Failed to create temp file.")

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{temp.Name()}
	myLog, _ := cfg.Build()
	for i := 0; i < b.N; i++ {
		myLog.Sugar().Infow(
			"random int bench",
			"random", rand.Intn(100),
		)
	}
}

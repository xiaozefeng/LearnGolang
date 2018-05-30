package logs

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "wlarus",
		"size": 10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// 通过日志语句重用字段
	// logrus.Entry返回自WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}

func init() {
	// 日志格式化为JSON而不是默认的ASCII
	log.SetFormatter(&log.JSONFormatter{})

	// 可以是一个文件
	log.SetOutput(os.Stdout)

	// 设置level
	log.SetLevel(log.WarnLevel)
}

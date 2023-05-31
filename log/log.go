package log

import "github.com/sirupsen/logrus"

// Log 创建一个新的日志记录器
var Log = logrus.New()

func init() {
	Log.SetLevel(logrus.InfoLevel)
	//Log.SetFormatter(&logrus.JSONFormatter{})
}

package helper

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"messageBot/conf"
	"os"
	"strings"
	"time"
)

const NOLOGFlAG string = "NOLOGFlAG"

func getLogLevel(loglevel string) (logrus.Level, bool) {
	switch loglevel {
	case "debug":
		return logrus.DebugLevel, true
	case "info":
		return logrus.InfoLevel, true
	case "warn":
		return logrus.WarnLevel, true
	case "error":
		return logrus.ErrorLevel, true
	case "fatal":
		return logrus.FatalLevel, true
	case "panic":
		return logrus.PanicLevel, true
	default:
		return 0, false
	}
}

func InitLog() {
	loglevel := conf.Conf.LogLevel
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	logrus.SetOutput(os.Stdout)
	if level, exist := getLogLevel(loglevel); exist {
		logrus.SetLevel(level)
	} else {
		if GetEnv() == ENV_DEV {
			logrus.SetLevel(logrus.DebugLevel)
		} else if GetEnv() == ENV_TEST {
			logrus.SetLevel(logrus.DebugLevel)
		} else if GetEnv() == ENV_PRODUCT {
			logrus.SetLevel(logrus.InfoLevel)
		}
	}

}

func Logger(ctx *gin.Context) *logrus.Entry {
	entry := logrus.NewEntry(logrus.StandardLogger()).WithFields(logrus.Fields{
		"requestId": GetRequestId(ctx),
	})

	return entry
}

func DebugLogger(ctx *gin.Context, value ...interface{}) {
	if GetNoLogFlag(ctx) == true {
		return
	}
	if ctx != nil {
		Logger(ctx).Debug(value...)
	} else {
		logrus.Debug(value...)
	}
}

func DebugfLogger(ctx *gin.Context, format string, value ...interface{}) {
	if GetNoLogFlag(ctx) == true {
		return
	}
	if ctx != nil {
		Logger(ctx).Debugf(format, value...)
	} else {
		logrus.Debugf(format, value...)
	}
}

func InfoLogger(ctx *gin.Context, value ...interface{}) {
	if GetNoLogFlag(ctx) == true {
		return
	}
	if ctx != nil {
		Logger(ctx).Info(value...)
	} else {
		logrus.Info(value...)
	}
}

func InfofLogger(ctx *gin.Context, format string, value ...interface{}) {
	if GetNoLogFlag(ctx) == true {
		return
	}
	if ctx != nil {
		Logger(ctx).Infof(format, value...)
	} else {
		logrus.Infof(format, value...)
	}
}

func WarnLogger(ctx *gin.Context, value ...interface{}) {
	if GetNoLogFlag(ctx) == true {
		return
	}
	if ctx != nil {
		Logger(ctx).Warn(value...)
	} else {
		logrus.Warn(value...)
	}
}

func WarnfLogger(ctx *gin.Context, format string, value ...interface{}) {
	if GetNoLogFlag(ctx) == true {
		return
	}
	if ctx != nil {
		Logger(ctx).Warnf(format, value...)
	} else {
		logrus.Warnf(format, value...)
	}
}

func ErrorLogger(ctx *gin.Context, value ...interface{}) {
	if ctx != nil {
		Logger(ctx).Error(value...)
	} else {
		logrus.Error(value...)
	}
}

func ErrorfLogger(ctx *gin.Context, format string, value ...interface{}) {
	if ctx != nil {
		Logger(ctx).Errorf(format, value...)
	} else {
		logrus.Errorf(format, value...)
	}
}

func PanicLogger(ctx *gin.Context, value ...interface{}) {
	if ctx != nil {
		Logger(ctx).Panic(value...)
	} else {
		logrus.Panic(value...)
	}
}

func PanicfLogger(ctx *gin.Context, format string, value ...interface{}) {
	if ctx != nil {
		Logger(ctx).Panicf(format, value...)
	} else {
		logrus.Panicf(format, value...)
	}
}

func SetNoLogFlag(ctx *gin.Context) {
	ctx.Set(NOLOGFlAG, true)
}

func GetNoLogFlag(ctx *gin.Context) bool {
	if ctx == nil {
		return false
	}
	flag, ok := ctx.Get(NOLOGFlAG)
	if ok == true && flag == true {
		return true
	} else {
		return false
	}
}

func StackLogger(ctx *gin.Context, err error) {
	if !strings.Contains(fmt.Sprintf("%+v", err), "\n") {
		return
	}

	var info []byte
	if ctx != nil {
		info, _ = json.Marshal(map[string]interface{}{"time": time.Now().Format("2006-01-02 15:04:05"), "level": "error", "module": "errorstack", "requestId": GetRequestId(ctx)})
	} else {
		info, _ = json.Marshal(map[string]interface{}{"time": time.Now().Format("2006-01-02 15:04:05"), "level": "error", "module": "errorstack"})
	}

	fmt.Printf("%s\n-------------------stack-start-------------------\n%+v\n-------------------stack-end-------------------\n", string(info), err)
}

func GetRequestId(ctx *gin.Context) string {
	if ctx == nil {
		return ""
	}
	requestId, exist := ctx.Get("requestId")
	if exist {
		return requestId.(string)
	}
	return ""
}

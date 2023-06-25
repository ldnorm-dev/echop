package echop

import (
	"github.com/labstack/echo/v4"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	syslog "log"
	"os"
	"path/filepath"
)

type (
	LoggerFields struct {
		Msg    string
		Fields []zap.Field
		C      echo.Context
	}
)

// InitLogger initializes the logger.
// default logger is initialized if l is nil.
func InitLogger(l *zap.Logger) {
	if l == nil {
		core := zapcore.NewCore(getEncoder(), getLogWriter(), zapcore.DebugLevel)
		Logger = zap.New(core)
	} else {
		Logger = l
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// GetAppRunPath returns the path where the app is running.
func GetAppRunPath() (runPath string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	runPath = filepath.Dir(ex)
	if runPath[len(runPath)-1:] != "/" {
		runPath = runPath + "/"
	}
	return
}

func getLogWriter() zapcore.WriteSyncer {
	runPath := GetAppRunPath()
	if _, err := os.Stat(runPath + ".log"); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(runPath+".log", 0755); err != nil {
				syslog.Fatal(err)
			}
		}
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   runPath + ".log/" + AppName + ".log",
		MaxSize:    10,
		MaxBackups: 90,
		MaxAge:     90,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func log(logInfo LoggerFields) (string, []zap.Field) {
	if Logger == nil {
		InitLogger(nil)
	}
	fields := append(logInfo.Fields)
	if logInfo.C != nil {
		fields = append(fields, zap.String(RequestIDConfig.TargetHeader, GetRequestID(logInfo.C)))
	}
	return logInfo.Msg, fields
}

func LogInfoWithContext(c echo.Context, msg string, fields ...zap.Field) {
	msg, newFields := log(LoggerFields{
		Msg:    msg,
		Fields: fields,
		C:      c,
	})
	Logger.Info(msg, newFields...)
}

func LogErrorWithContext(c echo.Context, msg string, fields ...zap.Field) {
	msg, newFields := log(LoggerFields{
		Msg:    msg,
		Fields: fields,
		C:      c,
	})
	Logger.Error(msg, newFields...)
}

func LogWarnWithContext(c echo.Context, msg string, fields ...zap.Field) {
	msg, newFields := log(LoggerFields{
		Msg:    msg,
		Fields: fields,
		C:      c,
	})
	Logger.Warn(msg, newFields...)
}

func LogDebugWithContext(c echo.Context, msg string, fields ...zap.Field) {
	msg, newFields := log(LoggerFields{
		Msg:    msg,
		Fields: fields,
		C:      c,
	})
	Logger.Debug(msg, newFields...)
}

func LogInfo(msg string, fields ...zap.Field) {
	LogInfoWithContext(nil, msg, fields...)
}

func LogError(msg string, fields ...zap.Field) {
	LogErrorWithContext(nil, msg, fields...)
}

func LogWarn(msg string, fields ...zap.Field) {
	LogWarnWithContext(nil, msg, fields...)
}

func LogDebug(msg string, fields ...zap.Field) {
	LogDebugWithContext(nil, msg, fields...)
}

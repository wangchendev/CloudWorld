package logs

import (
    "go.uber.org/zap"
)

var defaultLogger *zap.SugaredLogger

func init()  {
    logger, _ := zap.NewDevelopment(zap.AddCallerSkip(1))
    defaultLogger = logger.Sugar()
}

func Debug(template string, args ...interface{}) {
    defaultLogger.Debugf(template, args)
}

func Info(template string, args ...interface{}) {
    defaultLogger.Infof(template, args)
}

func Warn(template string, args ...interface{}) {
    defaultLogger.Warnf(template, args)
}

func Error(template string, args ...interface{}) {
    defaultLogger.Errorf(template, args)
}

func DPanic(template string, args ...interface{}) {
    defaultLogger.DPanicf(template, args)
}

func Panic(template string, args ...interface{}) {
    defaultLogger.Panicf(template, args)
}

func Flush() {
    _ = defaultLogger.Sync()
}
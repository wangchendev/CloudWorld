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
    if len(args) == 0 {
        defaultLogger.Debug(template)
    } else {
        defaultLogger.Debugf(template, args)
    }
}

func Info(template string, args ...interface{}) {
    if len(args) == 0 {
        defaultLogger.Info(template)
    } else {
        defaultLogger.Infof(template, args)
    }
}

func Warn(template string, args ...interface{}) {
    if len(args) == 0 {
        defaultLogger.Warn(template)
    } else {
        defaultLogger.Warnf(template, args)
    }
}

func Error(template string, args ...interface{}) {
    if len(args) == 0 {
        defaultLogger.Error(template)
    } else {
        defaultLogger.Errorf(template, args)
    }
}

func DPanic(template string, args ...interface{}) {
    if len(args) == 0 {
        defaultLogger.DPanic(template)
    } else {
        defaultLogger.DPanicf(template, args)
    }
}

func Panic(template string, args ...interface{}) {
    if len(args) == 0 {
        defaultLogger.Panic(template)
    } else {
        defaultLogger.Panicf(template, args)
    }
}

func Flush() {
    _ = defaultLogger.Sync()
}
package mysql

import (
	"context"
	"github.com/carefreex-io/logger"
	gormLog "gorm.io/gorm/logger"
	"time"
)

type GormLog struct {
	LogLevel gormLog.LogLevel
}

func New() gormLog.Interface {
	return &GormLog{}
}

// LogMode log mode
func (l *GormLog) LogMode(level gormLog.LogLevel) gormLog.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info
func (l GormLog) Info(ctx context.Context, msg string, data ...interface{}) {
	tag := "GormLog.Info"
	if l.LogLevel >= gormLog.Info {
		logger.Infox(ctx, tag, append([]interface{}{msg}, data...)...)
	}
}

// Warn print warn messages
func (l GormLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	//if l.LogLevel >= gormLog.Warn {
	//	l.Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	//}
}

// Error print error messages
func (l GormLog) Error(ctx context.Context, msg string, data ...interface{}) {
	//if l.LogLevel >= Error {
	//	l.Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	//}
}

// Trace print sql message
func (l GormLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	//if l.LogLevel <= Silent {
	//	return
	//}
	//
	//elapsed := time.Since(begin)
	//switch {
	//case err != nil && l.LogLevel >= Error && (!errors.Is(err, ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
	//	sql, rows := fc()
	//	if rows == -1 {
	//		l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
	//	} else {
	//		l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	//	}
	//case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= Warn:
	//	sql, rows := fc()
	//	slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
	//	if rows == -1 {
	//		l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
	//	} else {
	//		l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	//	}
	//case l.LogLevel == Info:
	//	sql, rows := fc()
	//	if rows == -1 {
	//		l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
	//	} else {
	//		l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
	//	}
	//}
}

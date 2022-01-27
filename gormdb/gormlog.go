package gormdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/carefreex-io/logger"
	gormLog "gorm.io/gorm/logger"
	"time"
)

type GormLog struct {
	gormLog.Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func New(config gormLog.Config) gormLog.Interface {
	var (
		infoStr      = "[info] "
		warnStr      = "[warn] "
		errStr       = "[error] "
		traceStr     = "[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		infoStr = gormLog.Green + "\n" + gormLog.Reset + gormLog.Green + "[info] " + gormLog.Reset
		warnStr = gormLog.BlueBold + "\n" + gormLog.Reset + gormLog.Magenta + "[warn] " + gormLog.Reset
		errStr = gormLog.Magenta + "\n" + gormLog.Reset + gormLog.Red + "[error] " + gormLog.Reset
		traceStr = gormLog.Green + "\n" + gormLog.Reset + gormLog.Yellow + "[%.3fms] " + gormLog.BlueBold + "[rows:%v]" + gormLog.Reset + " %s"
		traceWarnStr = gormLog.Green + "" + gormLog.Yellow + "%s\n" + gormLog.Reset + gormLog.RedBold + "[%.3fms] " + gormLog.Yellow + "[rows:%v]" + gormLog.Magenta + " %s" + gormLog.Reset
		traceErrStr = gormLog.RedBold + "" + gormLog.MagentaBold + "%s\n" + gormLog.Reset + gormLog.Yellow + "[%.3fms] " + gormLog.BlueBold + "[rows:%v]" + gormLog.Reset + " %s"
	}

	return &GormLog{
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

// LogMode log mode
func (l *GormLog) LogMode(level gormLog.LogLevel) gormLog.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info
func (l GormLog) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLog.Info {
		logger.InfofX(ctx, l.infoStr+msg, data...)
	}
}

// Warn print warn messages
func (l GormLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLog.Warn {
		logger.WarnfX(ctx, l.warnStr+msg, data...)
	}
}

// Error print error messages
func (l GormLog) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLog.Error {
		logger.ErrorfX(ctx, l.errStr+msg, data...)
	}
}

// Trace print sql message
func (l GormLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= gormLog.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormLog.Error && (!errors.Is(err, gormLog.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			logger.ErrorfX(ctx, l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			logger.ErrorfX(ctx, l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= gormLog.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			logger.WarnfX(ctx, l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			logger.WarnfX(ctx, l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == gormLog.Info:
		sql, rows := fc()
		if rows == -1 {
			logger.InfofX(ctx, l.traceStr, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			logger.InfofX(ctx, l.traceStr, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}

package xormdb

import (
	"fmt"
	"github.com/xormplus/xorm/log"
)

type XormLogCtx struct {
	logger XormLog
}

func NewXormLogCtx(logger XormLog) log.ContextLogger {
	return &XormLogCtx{
		logger: logger,
	}
}

// BeforeSQL implements ContextLogger
func (l *XormLogCtx) BeforeSQL(ctx log.LogContext) {
	l.logger.ctx = ctx.Ctx
}

// AfterSQL implements ContextLogger
func (l *XormLogCtx) AfterSQL(ctx log.LogContext) {
	var sessionPart string
	v := ctx.Ctx.Value(log.SessionIDKey)
	if key, ok := v.(string); ok {
		sessionPart = fmt.Sprintf(" [%s]", key)
	}
	if ctx.ExecuteTime > 0 {
		l.logger.Infof("[SQL]%s %s %v - %v", sessionPart, ctx.SQL, ctx.Args, ctx.ExecuteTime)
	} else {
		l.logger.Infof("[SQL]%s %s %v", sessionPart, ctx.SQL, ctx.Args)
	}
}

// Debugf implements ContextLogger
func (l *XormLogCtx) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

// Errorf implements ContextLogger
func (l *XormLogCtx) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

// Infof implements ContextLogger
func (l *XormLogCtx) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

// Warnf implements ContextLogger
func (l *XormLogCtx) Warnf(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

// Level implements ContextLogger
func (l *XormLogCtx) Level() log.LogLevel {
	return l.logger.Level()
}

// SetLevel implements ContextLogger
func (l *XormLogCtx) SetLevel(lv log.LogLevel) {
	l.logger.SetLevel(lv)
}

// ShowSQL implements ContextLogger
func (l *XormLogCtx) ShowSQL(show ...bool) {
	l.logger.ShowSQL(show...)
}

// IsShowSQL implements ContextLogger
func (l *XormLogCtx) IsShowSQL() bool {
	return l.logger.IsShowSQL()
}

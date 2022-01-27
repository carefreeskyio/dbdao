package xormdb

import (
	"context"
	"github.com/carefreex-io/logger"
	"github.com/xormplus/xorm/log"
)

type XormLog struct {
	ctx     context.Context
	level   log.LogLevel
	showSQL bool
}

// Error implement ILogger
func (s *XormLog) Error(v ...interface{}) {
	if s.level <= log.LOG_ERR {
		logger.ErrorX(s.ctx, v...)
	}
	return
}

// Errorf implement ILogger
func (s *XormLog) Errorf(format string, v ...interface{}) {
	if s.level <= log.LOG_ERR {
		logger.ErrorfX(s.ctx, format, v...)
	}
	return
}

// Debug implement ILogger
func (s *XormLog) Debug(v ...interface{}) {
	if s.level <= log.LOG_DEBUG {
		logger.DebugX(s.ctx, v...)
	}
	return
}

// Debugf implement ILogger
func (s *XormLog) Debugf(format string, v ...interface{}) {
	if s.level <= log.LOG_DEBUG {
		logger.DebugfX(s.ctx, format, v...)
	}
	return
}

// Info implement ILogger
func (s *XormLog) Info(v ...interface{}) {
	if s.level <= log.LOG_INFO {
		logger.InfoX(s.ctx, v...)
	}
	return
}

// Infof implement ILogger
func (s *XormLog) Infof(format string, v ...interface{}) {
	if s.level <= log.LOG_INFO {
		logger.InfofX(s.ctx, format, v...)
	}
	return
}

// Warn implement ILogger
func (s *XormLog) Warn(v ...interface{}) {
	if s.level <= log.LOG_WARNING {
		logger.WarnX(s.ctx, v...)
	}
	return
}

// Warnf implement ILogger
func (s *XormLog) Warnf(format string, v ...interface{}) {
	if s.level <= log.LOG_WARNING {
		logger.WarnfX(s.ctx, format, v...)
	}
	return
}

// Level implement ILogger
func (s *XormLog) Level() log.LogLevel {
	return s.level
}

// SetLevel implement ILogger
func (s *XormLog) SetLevel(l log.LogLevel) {
	s.level = l
	return
}

// ShowSQL implement ILogger
func (s *XormLog) ShowSQL(show ...bool) {
	if len(show) == 0 {
		s.showSQL = true
		return
	}
	s.showSQL = show[0]
}

// IsShowSQL implement ILogger
func (s *XormLog) IsShowSQL() bool {
	return s.showSQL
}

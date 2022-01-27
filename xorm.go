// +build xorm

package dbdao

import (
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/log"
	"github.com/xormplus/xorm/names"
	"sync"
	"time"
)

var (
	Read  *xorm.Engine
	Write *xorm.Engine
)

type Log struct {
	LogLevel log.LogLevel
	ShowSql  bool
}

type BaseOptions struct {
	ReadDns     string
	WriteDns    string
	Logger      Log
	MaxIdleConn int
	MaxOpenConn int
	MaxLifetime time.Duration
	TablePrefix string
}

type CustomOptions struct {
}

var (
	dbOnce               sync.Once
	baseOptions          *BaseOptions
	DefaultCustomOptions = &CustomOptions{}
)

func initBaseOptions() {
	baseOptions = &BaseOptions{
		ReadDns:     config.GetString("Mysql.Read"),
		WriteDns:    config.GetString("Mysql.Write"),
		MaxIdleConn: config.GetInt("MysqlConf.MaxIdleConn"),
		MaxOpenConn: config.GetInt("MysqlConf.MaxOpenConn"),
		MaxLifetime: config.GetDuration("MysqlConf.MaxLifetime") * time.Second,
		Logger: Log{
			LogLevel: log.LOG_INFO,
			ShowSql:  config.GetBool("MysqlConf.Log.ShowSql"),
		},
		TablePrefix: config.GetString("MysqlConf.TablePrefix"),
	}
}

func InitDB() (err error) {
	dbOnce.Do(func() {
		initBaseOptions()

		if baseOptions.ReadDns == "" && baseOptions.WriteDns == "" {
			return
		}

		if Read, err = xorm.NewEngine("mysql", baseOptions.ReadDns); err != nil {
			logger.Errorf("read xorm.NewEngine failed: err=%v", err)
			return
		}
		setConfig(Read)

		if Write, err = xorm.NewEngine("mysql", baseOptions.WriteDns); err != nil {
			logger.Errorf("write xorm.NewEngine failed: err=%v", err)
			return
		}
		setConfig(Write)
	})

	return err
}

func setConfig(db *xorm.Engine) {
	db.SetMaxIdleConns(baseOptions.MaxIdleConn)
	db.SetMaxOpenConns(baseOptions.MaxOpenConn)
	db.SetConnMaxLifetime(baseOptions.MaxLifetime)
	xormLog := XormLog{}
	xormLogCtx := NewXormLogCtx(xormLog)
	xormLogCtx.SetLevel(baseOptions.Logger.LogLevel)
	xormLogCtx.ShowSQL(baseOptions.Logger.ShowSql)
	db.SetLogger(xormLogCtx)
	tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, baseOptions.TablePrefix)
	db.SetTableMapper(tbMapper)
}

func NewReadSession() (db *xorm.Session) {
	return Read.NewSession()
}

func NewWriteSession() (db *xorm.Session) {
	return Write.NewSession()
}

package mysql
//
//import (
//	"github.com/carefreex-io/config"
//	"github.com/carefreex-io/logger"
//	"github.com/xormplus/xorm"
//	"github.com/xormplus/xorm/log"
//	"os"
//	"sync"
//)
//
//type DB struct {
//	Read  *xorm.Engine
//	Write *xorm.Engine
//}
//
//type BaseOptions struct {
//	ReadDns  string
//	WriteDns string
//}
//
//type CustomOptions struct {
//	ShowSql bool
//}
//
//var (
//	Db                   *DB
//	dbOnce               sync.Once
//	baseOptions          *BaseOptions
//	DefaultCustomOptions = &CustomOptions{}
//)
//
//func initBaseOptions() {
//	baseOptions = &BaseOptions{
//		ReadDns:  config.GetString("Mysql.Read"),
//		WriteDns: config.GetString("Mysql.Write"),
//	}
//}
//
//func InitDB(customOptions *CustomOptions) {
//	initBaseOptions()
//
//	var err error
//
//	if Db == nil {
//		dbOnce.Do(func() {
//			if Db.Read, err = xorm.NewEngine("mysql", baseOptions.ReadDns); err != nil {
//				logger.Errorf("xorm.NewEngine failed: err=%v", err)
//				return
//			}
//			logger := log.NewSimpleLogger()
//			Db.Read.SetLogger(log.NewLoggerAdapter(logger))
//			Db.Read.SetMaxIdleConns()
//
//			if Db.Write, err = xorm.NewEngine("mysql", baseOptions.WriteDns); err != nil {
//				logger.Errorf("xorm.NewEngine failed: err=%v", err)
//				return
//			}
//		})
//	}
//}
//
//func NewReadSession() (db *xorm.Session) {
//	return Db.Read.NewSession()
//}
//
//func NewWriteSession() (db *xorm.Session) {
//	return Db.Write.NewSession()
//}

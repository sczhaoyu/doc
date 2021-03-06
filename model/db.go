package model

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

const (
	SQL_NUM     = 150 //SQL批处理条数
	MAX_CLIENT  = 400 //最大链接个数
	INIT_CLIENT = 10  //初始化链接个数
)

var (
	DocDB *xorm.Engine //数据库

)

func init() {
	//====================================================================
	url := "root:root_sjht!@#$133@tcp(101.200.215.133:3306)/"
	if os.Getenv("GO_DEV") == "2" {
		url = "root:root_sjht!@#$133@tcp(localhost:3306)/"
	}
	DocDB, _ = xorm.NewEngine("mysql", url+"doc?charset=utf8")
	if os.Getenv("GO_DEV") == "1" {
		DocDB.ShowSQL = true

	}
	DocDB.SetMaxIdleConns(INIT_CLIENT)
	DocDB.SetMaxOpenConns(MAX_CLIENT)
	//====================================================================

}
func NoData(b bool) error {

	if b {
		return nil
	}
	return errors.New("empty")
}

//错误消息定义
func NoDataMsg(b bool, msg string) error {
	if b {
		return nil
	}
	return errors.New(msg)
}

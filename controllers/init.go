package controllers

import (
	"github.com/keep94/ramstore"
	"net/http"
)

var session *ramstore.RAMStore
var loginSessionName string

func init() {
	loginSessionName = "loginInfo"       //登录session名称
	session = ramstore.NewRAMStore(3600) //创建session
}

//设置session
func SetSession(r *http.Request, w http.ResponseWriter, key string, val interface{}) {
	m, _ := session.New(r, key)
	m.Values[key] = val
	m.Save(r, w)
}

//获取session
func GetSession(r *http.Request, key string) interface{} {
	s, _ := session.Get(r, key)
	a := s.Values[key]
	if a == nil {
		return nil
	}
	return a
}

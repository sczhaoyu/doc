package controllers

import (
	"github.com/sczhaoyu/doc/model"
	"net/http"
	"strconv"
)

//获取全部项目信息
func findProject(w http.ResponseWriter, r *http.Request) {
	ret, err := model.FindProject()
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	w.Write(ToJson(ret))
}

//获取项目版本
func findProjectVersion(w http.ResponseWriter, r *http.Request) {
	projectId, _ := strconv.ParseInt(r.FormValue("projectId"), 10, 64)
	ret, err := model.FindVersion(projectId)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	w.Write(ToJson(ret))
}

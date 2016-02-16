package controllers

import (
	"errors"
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

//保存项目版本
func saveProjectVersion(w http.ResponseWriter, r *http.Request) {
	//版本名称
	version := r.FormValue("version")
	//所属项目
	projectId, _ := strconv.ParseInt(r.FormValue("projectId"), 10, 64)
	if version == "" {
		w.Write(ToJson(errors.New("版本名称不能为空！")))
		return
	}
	var v model.Version
	v.ProjectId = projectId
	v.Version = version
	w.Write(ToJson(v.Save()))
}
func cloneProjectVersion(w http.ResponseWriter, r *http.Request) {
	projectId, _ := strconv.ParseInt(r.FormValue("projectId"), 10, 64)
	oldVersionId, _ := strconv.ParseInt(r.FormValue("oldVersionId"), 10, 64)
	newVersionId, _ := strconv.ParseInt(r.FormValue("newVersionId"), 10, 64)
	if oldVersionId == newVersionId || oldVersionId == 0 || newVersionId == 0 {
		w.Write(ToJson(errors.New("版本号ID不能为空并且版本不能相同！")))
		return
	}
	//开始复制版本信息
	err := model.CopyVersion(projectId, oldVersionId, newVersionId)
	w.Write(ToJson(err))

}

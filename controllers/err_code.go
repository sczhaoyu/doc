package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/sczhaoyu/doc/model"
	"net/http"
	"strconv"
)

func saveErrCode(w http.ResponseWriter, r *http.Request) {
	j := r.FormValue("json")
	var c model.ErrCode
	err := json.Unmarshal([]byte(j), &c)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	if c.Id > 0 {
		err = c.Update()
		if err == nil {
			model.AddUpdateLog(fmt.Sprintf("修改错误代码【%s:%s】", c.Code, c.DescriptionText), c.ProjectId, c.VersionId)
		}
	} else {
		err = c.Save()
		if err == nil {
			model.AddUpdateLog(fmt.Sprintf("增加错误代码【%s:%s】", c.Code, c.DescriptionText), c.ProjectId, c.VersionId)
		}
	}

	w.Write(ToJson(err))
}
func updateErrCode(w http.ResponseWriter, r *http.Request) {
	j := r.FormValue("json")
	var c model.ErrCode
	err := json.Unmarshal([]byte(j), &c)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	err = c.Update()
	if err == nil {
		model.AddUpdateLog(fmt.Sprintf("修改错误代码【%s:%s】", c.Code, c.DescriptionText), c.ProjectId, c.VersionId)
	}
	w.Write(ToJson(err))
}
func getErrCodeAll(w http.ResponseWriter, r *http.Request) {
	projectId, _ := strconv.ParseInt(r.FormValue("projectId"), 10, 64)
	ret, err := model.FindErrCode(projectId)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	w.Write(ToJson(ret))
}
func deleteErrCode(w http.ResponseWriter, r *http.Request) {
	eid, _ := strconv.ParseInt(r.FormValue("eid"), 10, 64)
	ret, err := model.GetErrCodeById(eid)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	err = ret.Delete()
	model.AddUpdateLog(fmt.Sprintf("删除错误代码【%s】", ret.Code), ret.ProjectId, ret.VersionId)
	w.Write(ToJson(err))
}

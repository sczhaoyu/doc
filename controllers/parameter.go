package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/sczhaoyu/doc/model"
	"net/http"
	"strconv"
)

//更新参数
func updateParameter(w http.ResponseWriter, r *http.Request) {
	var p model.Parameters
	j := r.FormValue("json")
	err := json.Unmarshal([]byte(j), &p)
	if err != nil {
		w.Write(ToJson(err))
		return
	}

	err = p.Update()
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	doc, err := model.GetDocById(p.DocId)
	if err == nil {
		pt := "请求参数"
		if p.PrmType == 1 {
			pt = "响应参数"
		}
		model.AddUpdateLog(fmt.Sprintf("文档【%s】【%s】参数被修改,参数编号【%s】", doc.Name, pt, p.SerialNumber), doc.ProjectId, doc.VersionId)
	}
	w.Write(ToJson("success"))

}

//更新参数
func deleteParameter(w http.ResponseWriter, r *http.Request) {
	var p model.Parameters
	p.Id, _ = strconv.ParseInt(r.FormValue("pid"), 10, 64)
	err := p.Delete()
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	w.Write(ToJson("success"))

}

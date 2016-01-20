package controllers

import (
	"encoding/json"
	"github.com/sczhaoyu/doc/model"
	"net/http"
	"strconv"
)

//获取全部父目录
func getCatalogueAll(w http.ResponseWriter, r *http.Request) {
	projectId, _ := strconv.ParseInt(r.FormValue("projectId"), 10, 64)
	versionId, _ := strconv.ParseInt(r.FormValue("versionId"), 10, 64)
	ret, err := model.FindCatalogueAllParent(projectId, versionId)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	w.Write(ToJson(ret))
}

//获取父目录的子目录
func getCatalogueChild(w http.ResponseWriter, r *http.Request) {
	catalogueId, _ := strconv.ParseInt(r.FormValue("catalogueId"), 10, 64)
	projectId, _ := strconv.ParseInt(r.FormValue("projectId"), 10, 64)
	versionId, _ := strconv.ParseInt(r.FormValue("versionId"), 10, 64)
	ret, err := model.FindFindCatalogueByParent(projectId, versionId, catalogueId)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	w.Write(ToJson(ret))
}
func findDoc(w http.ResponseWriter, r *http.Request) {
	catalogueId, _ := strconv.ParseInt(r.FormValue("catalogueId"), 10, 64)
	ret, err := model.FindDoc(catalogueId)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	w.Write(ToJson(ret))
}

//获取文档详情
func getCatalogueDoc(w http.ResponseWriter, r *http.Request) {
	docId, _ := strconv.ParseInt(r.FormValue("docId"), 10, 64)
	doc, err := model.GetDocById(docId)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	p, _ := model.FindParametersByDocId(doc.Id)
	ret := struct {
		Doc        model.Doc          `json:"doc"`
		Parameters []model.Parameters `json:"parameters"`
	}{
		Doc:        *doc,
		Parameters: p,
	}
	w.Write(ToJson(ret))
}

type SubmitJson struct {
	Doc           model.Doc          `json:"doc"`
	RspParameters []model.Parameters `json:"rspParameters"`
	ReqParameters []model.Parameters `json:"reqParameters"`
}

func catalogueSubmit(w http.ResponseWriter, r *http.Request) {
	j := r.FormValue("json")
	var p SubmitJson
	err := json.Unmarshal([]byte(j), &p)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	//保存文档
	if p.Doc.Id > 0 {
		err = p.Doc.Update()
		if err == nil {
			//添加操作日志
			model.AddUpdateLog("修改文档【"+p.Doc.Name+":"+p.Doc.SerialNumber+"】", p.Doc.ProjectId, p.Doc.VersionId)
			err = model.DeleteParameters(p.Doc.Id)
		}
	} else {
		model.AddUpdateLog("增加文档【"+p.Doc.Name+":"+p.Doc.SerialNumber+"】", p.Doc.ProjectId, p.Doc.VersionId)
		err = p.Doc.Save()
	}

	if err != nil {
		w.Write(ToJson(err))
		return
	}
	var prm []model.Parameters = make([]model.Parameters, 0, len(p.ReqParameters)+len(p.RspParameters))
	for i := 0; i < len(p.RspParameters); i++ {
		p.RspParameters[i].Id = 0
		p.RspParameters[i].PrmType = 1
		p.RspParameters[i].DocId = p.Doc.Id
		p.RspParameters[i].ProjectId = p.Doc.ProjectId
		p.RspParameters[i].VersionId = p.Doc.VersionId
		prm = append(prm, p.RspParameters[i])
	}
	for i := 0; i < len(p.ReqParameters); i++ {
		p.ReqParameters[i].Id = 0
		p.ReqParameters[i].PrmType = 0
		p.ReqParameters[i].DocId = p.Doc.Id
		p.ReqParameters[i].ProjectId = p.Doc.ProjectId
		p.ReqParameters[i].VersionId = p.Doc.VersionId
		prm = append(prm, p.ReqParameters[i])
	}
	if len(prm) > 0 {
		err = model.SaveParameters(prm)
	}
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	w.Write(ToJson("保存成功！"))
}

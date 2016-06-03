package controllers

import (
	"encoding/json"
	"errors"
	"github.com/sczhaoyu/doc/model"
	"net/http"
	"strconv"
	"strings"
)

func updateVersion(w http.ResponseWriter, r *http.Request) {
	versionId, _ := strconv.ParseInt(r.FormValue("versionId"), 10, 64)
	v, err := model.GetVersionById(versionId)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	v.Version = r.FormValue("version")
	w.Write(ToJson(v.Update()))

}
func saveCatalogue(w http.ResponseWriter, r *http.Request) {
	projectId, _ := strconv.ParseInt(r.FormValue("projectId"), 10, 64)
	versionId, _ := strconv.ParseInt(r.FormValue("versionId"), 10, 64)
	var c model.Catalogue
	c.Name = strings.Trim(r.FormValue("name"), " ")
	c.ProjectId = projectId
	c.SerialNumber = strings.Trim(r.FormValue("serialNumber"), " ")
	c.VersionId = versionId
	if c.Name == "" {
		w.Write(ToJson(errors.New("目录名称不能为空！")))
		return
	}
	if c.SerialNumber == "" {
		w.Write(ToJson(errors.New("目录序号不能为空！")))
		return
	}
	w.Write(ToJson(c.Save()))

}
func updateCatalogue(w http.ResponseWriter, r *http.Request) {
	catalogueId, _ := strconv.ParseInt(r.FormValue("catalogueId"), 10, 64)
	c, err := model.GetCatalogueById(catalogueId)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	c.Name = r.FormValue("name")
	c.SerialNumber = r.FormValue("serialNumber")
	if c.Name == "" || c.SerialNumber == "" {
		w.Write(ToJson(errors.New("目录编号和目录名称不能为空！")))
		return
	}
	w.Write(ToJson(c.Update()))
}
func deleteCatalogue(w http.ResponseWriter, r *http.Request) {
	catalogueId, _ := strconv.ParseInt(r.FormValue("catalogueId"), 10, 64)
	w.Write(ToJson(model.DeleteCatalogue(catalogueId)))
}

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
func copyCatalogueDoc(w http.ResponseWriter, r *http.Request) {
	docId, _ := strconv.ParseInt(r.FormValue("docId"), 10, 64)
	catalogueId, _ := strconv.ParseInt(r.FormValue("catalogueId"), 10, 64)
	versionId, _ := strconv.ParseInt(r.FormValue("versionId"), 10, 64)
	projectId, _ := strconv.ParseInt(r.FormValue("projectId"), 10, 64)
	name := r.FormValue("name")
	serialNumber := r.FormValue("serialNumber")
	if catalogueId <= 0 || versionId <= 0 || projectId <= 0 {
		w.Write(ToJson(errors.New("参数错误！")))
		return
	}
	w.Write(ToJson(model.CopyDoc(docId, catalogueId, projectId, versionId, name, serialNumber)))

}

func deleteCatalogueDoc(w http.ResponseWriter, r *http.Request) {
	docId, _ := strconv.ParseInt(r.FormValue("docId"), 10, 64)
	var doc model.Doc
	doc.Id = docId
	w.Write(ToJson(doc.Delete()))

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

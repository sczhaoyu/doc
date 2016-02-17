package controllers

import (
	"net/http"
)

func steupRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/catalogue/all", getCatalogueAll)
	http.HandleFunc("/catalogue/child", getCatalogueChild)
	http.HandleFunc("/catalogue/save", saveCatalogue)
	http.HandleFunc("/catalogue/find/doc", findDoc)
	http.HandleFunc("/catalogue/doc", getCatalogueDoc)
	http.HandleFunc("/catalogue/doc/copy", copyCatalogueDoc)
	http.HandleFunc("/catalogue/doc/delete", deleteCatalogueDoc)
	http.HandleFunc("/catalogue/submit", catalogueSubmit)
	http.HandleFunc("/parameter/update", updateParameter)
	http.HandleFunc("/parameter/delete", deleteParameter)
	http.HandleFunc("/err/code/save", saveErrCode)
	http.HandleFunc("/err/code/update", updateErrCode)
	http.HandleFunc("/err/code/all", getErrCodeAll)
	http.HandleFunc("/err/code/delete", deleteErrCode)
	http.HandleFunc("/explain/submit", explainSubmit)
	http.HandleFunc("/explain/find", explainFind)
	http.HandleFunc("/explain/delete", explainDelete)
	http.HandleFunc("/login_submit", loginSubmit)
	http.HandleFunc("/update/log/find", updateLogFind)
	http.HandleFunc("/project", findProject)
	http.HandleFunc("/project/version", findProjectVersion)
	//保存项目版本
	http.HandleFunc("/project/version/save", saveProjectVersion)
	//克隆版本
	http.HandleFunc("/project/version/clone", cloneProjectVersion)
}

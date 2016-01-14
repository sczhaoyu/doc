package controllers

import (
	"net/http"
)

func steupRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/catalogue/all", getCatalogueAll)
	http.HandleFunc("/catalogue/child", getCatalogueChild)
	http.HandleFunc("/catalogue/find/doc", findDoc)
	http.HandleFunc("/catalogue/doc", getCatalogueDoc)
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
}

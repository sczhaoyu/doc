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
}

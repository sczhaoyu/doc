package controllers

import (
	"github.com/sczhaoyu/doc/model"
	"net/http"
	"strconv"
)

func updateLogFind(w http.ResponseWriter, r *http.Request) {
	projectId, _ := strconv.ParseInt(r.FormValue("projectId"), 10, 64)
	ret, err := model.FindUpdateLog(projectId)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	w.Write(ToJson(ret))
}

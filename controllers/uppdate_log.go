package controllers

import (
	"github.com/sczhaoyu/doc/model"
	"net/http"
)

func updateLogFind(w http.ResponseWriter, r *http.Request) {
	ret, err := model.FindUpdateLog()
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	w.Write(ToJson(ret))
}

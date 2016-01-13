package controllers

import (
	"encoding/json"
	"github.com/sczhaoyu/doc/model"
	"net/http"
)

func saveErrCode(w http.ResponseWriter, r *http.Request) {
	j := r.FormValue("json")
	var c model.ErrCode
	err := json.Unmarshal([]byte(j), &c)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	err = c.Save()
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
	w.Write(ToJson(err))
}
func getErrCodeAll(w http.ResponseWriter, r *http.Request) {
	ret, err := model.FindErrCode()
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	w.Write(ToJson(ret))
}

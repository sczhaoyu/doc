package controllers

import (
	"encoding/json"
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
	} else {
		err = c.Save()
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
func deleteErrCode(w http.ResponseWriter, r *http.Request) {
	eid, _ := strconv.ParseInt(r.FormValue("eid"), 10, 64)
	var e model.ErrCode
	e.Id = eid
	err := e.Delete()
	w.Write(ToJson(err))
}

package controllers

import (
	"encoding/json"
	"github.com/sczhaoyu/doc/model"
	"net/http"
	"strconv"
	"time"
)

func explainSubmit(w http.ResponseWriter, r *http.Request) {
	j := r.FormValue("json")
	var e model.Explain
	err := json.Unmarshal([]byte(j), &e)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	if e.Id > 0 {
		err = e.Update()
	} else {
		e.CreatedAt = time.Now().Local()
		err = e.Save()
	}

	w.Write(ToJson(err))
}
func explainFind(w http.ResponseWriter, r *http.Request) {
	ret, err := model.FindExplain()
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	w.Write(ToJson(ret))
}

func explainDelete(w http.ResponseWriter, r *http.Request) {
	eid, _ := strconv.ParseInt(r.FormValue("eid"), 10, 64)
	var e model.Explain
	e.Id = eid
	err := e.Delete()
	w.Write(ToJson(err))
}

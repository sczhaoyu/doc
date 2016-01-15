package controllers

import (
	"encoding/json"
	"fmt"
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
		if err == nil {
			model.AddUpdateLog(fmt.Sprintf("修改文章【%s】", e.Title))
		}
	} else {
		e.CreatedAt = time.Now().Local()
		err = e.Save()
		if err == nil {
			model.AddUpdateLog(fmt.Sprintf("增加文章【%s】", e.Title))
		}
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
	ret, err := model.GetExplainById(eid)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	err = ret.Delete()
	if err == nil {
		model.AddUpdateLog(fmt.Sprintf("删除文章【%s】", ret.Title))
	}
	w.Write(ToJson(err))
}

package controllers

import (
	"github.com/sczhaoyu/doc/model"
	"net/http"
)

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	pwd := r.FormValue("password")
	ret, err := model.GetAccount(user, pwd)
	if err != nil {
		w.Write(ToJson(err))
		return
	}
	SetSession(r, w, loginSessionName, ret)
	w.Write(ToJson(ret))
}

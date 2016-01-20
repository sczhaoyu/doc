package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	sessionLogin := GetSession(r, loginSessionName)
	j := "null"
	if sessionLogin != nil {
		data, _ := json.Marshal(sessionLogin)
		j = string(data)
	}
	css := "http://cdn.centwaytech.com/doc/bundle.css"
	js := "http://cdn.centwaytech.com/doc/bundle.js"
	if os.Getenv("GO_DEV") == "1" {
		css = "http://localhost:3000/build/bundle.css"
		js = "http://localhost:3000/build/bundle.js"
	}
	tpl := `<!DOCTYPE html>
	<html>
	<head>
	    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
	    <title>首页</title>
	    <meta http-equiv="Content-Type" content="text/html; charset=utf8"/>
	    <meta name="apple-mobile-web-app-capable" content="yes"/>
	    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no"/>
		<link href="` + css + `" rel="stylesheet"/>
		<script>
		   var user=` + j + `;
		   var pro={projectId:0,name:""};
		   var version={versionId:0,version:""};
		   function dialog(b){
	 		  	layer.alert(b, {
	 		  		title:"提示",
				    skin: 'layui-layer-molv',
				    closeBtn: 0
				});
		   }
		</script>
	</head>
	<body>
	<div id="main"></div>
	<script src="http://cdn.centwaytech.com/jq.js"></script>
	<script src="http://cdn.centwaytech.com/layer/layer.js"></script>
	<script src="` + js + `"></script>
	</body>
	</html>
	`
	t, _ := template.New("tpl").Parse(tpl)
	t.Execute(w, nil)
}

package controllers

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl := `<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <title>首页</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf8"/>
    <meta name="apple-mobile-web-app-capable" content="yes"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no"/>
	<link href="http://localhost:3000/build/bundle.css" rel="stylesheet"/>
	<script>
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
<script src="http://localhost:3000/build/bundle.js"></script>
</body>
</html>
`
	t, _ := template.New("tpl").Parse(tpl)
	t.Execute(w, nil)
}

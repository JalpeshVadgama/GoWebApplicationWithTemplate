package main

import (
	"net/http"
	"text/template"
)

func main()  {
	http.HandleFunc("/",func(w http.ResponseWriter,req *http.Request){
		w.Header().Add("Cotent type", "text/html")
		tmpl,err:= template.New("test").Parse(doc)
		if err == nil{
			tmpl.Execute(w,nil)
		}
	})
	http.ListenAndServe(":8080",nil)
}

const doc = `
<!DOCTYPE html>
<html>
<head>
<title> Static Template</title>
</head>
<body>
<h1> Static template</h1>
</body>
</html>
`



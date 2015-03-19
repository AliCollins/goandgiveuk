// goandgive project goandgive.go
package goandgiveuk

import (
	"net/http"

	"html/template"
	//"appengine"
)

var templates = template.Must(template.ParseFiles(
	"templates/index.html",
	))

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	//c := appengine.NewContext(r)

	err := templates.ExecuteTemplate(w, "index.html", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

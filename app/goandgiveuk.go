// goandgive project goandgive.go
package goandgiveuk

import (
	"net/http"

	//	"appengine"
	"html/template"
)

var templates = template.Must(template.ParseFiles(
	"templates/index.html",
	"templates/contact.html",
	"templates/elements.html",
	"templates/generic.html",
))

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/contact.html", contact)
	http.HandleFunc("/elements.html", elements)
	http.HandleFunc("/generic.html", generic)
}

func root(w http.ResponseWriter, r *http.Request) {
	//	c := appengine.NewContext(r)
	//	var temp_site string = r.URL.String()
	//	c.Infof("Requested and truncated URL: %v", temp_site)

	//	if r.URL.String()[:1] == "/" {
	//		temp_site = temp_site[1:]
	//	}
	//	if temp_site == "" {
	//		temp_site = "index.html"
	//	}

	err := templates.ExecuteTemplate(w, "index.html", nil)
	//	err := templates.ExecuteTemplate(w, temp_site, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "contact.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func elements(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "elements.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func generic(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "generic.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

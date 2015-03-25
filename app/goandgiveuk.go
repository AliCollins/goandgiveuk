// goandgive project goandgive.go
package goandgiveuk

import (
	"html/template"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
)

var templates = template.Must(template.ParseFiles(
	"templates/index.html",
	"templates/contact.html",
	"templates/elements.html",
	"templates/generic.html",
))

type SignUpInfo struct {
	Email string
	IP    string
	Date  time.Time
}

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/contact.html", contact)
	http.HandleFunc("/elements.html", elements)
	http.HandleFunc("/generic.html", generic)
}

// Returns the key used for all signup entries
func signupKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "GoandGiveUK", "signup", 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	//	c := appengine.NewContext(r)

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	sign := SignUpInfo{
		Email: r.FormValue("email"),
		IP:    r.RemoteAddr,
		Date:  time.Now(),
	}
	// Use the same parent key on every signup to ensure each is in the same entity group.
	key := datastore.NewIncompleteKey(c, "SignUp", signupKey(c))
	_, err := datastore.Put(c, key, &sign)
	if err != nil {
		c.Infof("Failed to Signup: %v", sign.Email)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		c.Infof("Signup: %v", sign.Email)
	}

	http.Redirect(w, r, "/", http.StatusFound)
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

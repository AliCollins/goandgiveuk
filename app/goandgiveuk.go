// goandgive project goandgive.go
package goandgiveuk

import (
	"html/template"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/mail"
)

var templates = template.Must(template.ParseFiles(
	"templates/index.html",
	"templates/contact.html",
	"templates/elements.html",
	"templates/_header.html",
	"templates/_footer.html",
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
	http.HandleFunc("/contactSubmission", contactSubmission)
	http.HandleFunc("/elements.html", elements)
}

// Returns the key used for all signup entries
func signupKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "GoandGiveUK", "signup", 0, nil)
}

// Serve the main page using templates
func root(w http.ResponseWriter, r *http.Request) {

	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Sign-up method on root page
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

	// Email the details to admin
	// TODO only email based on number of entries - 1-10, 10s to 200, 100s to 1000, 250s beyond
	emailBody := sign.Email + "\nIP: " + sign.IP + "\n" + sign.Date.String()
	msg := &mail.Message{
		Sender: sign.Email,
		//To:		admin@goandgiveuk.appspot.com
		Subject: "GoAndGiveUK - Sign Up",
		Body:    emailBody,
	}
	// Send email to admins

	err = mail.SendToAdmins(c, msg)
	if err != nil {
		c.Infof("Failed to send sign-up email to Admin: %v", err)
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

// Contact Submission method on root page
func contactSubmission(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// Email the details to admin
	emailBody := "Name: " + r.FormValue("name") + "\nEmail: " + r.FormValue("email") +
		"\nSubject: " + r.FormValue("subject") + "\nMessage:\n" + r.FormValue("message")
	msg := &mail.Message{
		Sender: r.FormValue("email"),
		//To:		admin@goandgiveuk.appspot.com
		Subject: "GoAndGiveUK - Contact",
		Body:    emailBody,
	}
	// Send email to admins

	err := mail.SendToAdmins(c, msg)
	if err != nil {
		c.Infof("Failed to send contact email to Admin: %v", err)
	} else {
		c.Infof("Sent contact email to Admin: \n%v", emailBody)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func elements(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "elements.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

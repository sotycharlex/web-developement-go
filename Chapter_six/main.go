package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeView *view.View
var contactView *view.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	homeView, err = template.ParseFiles("home.gohtml", "layout/footer.gohtml")
	if err != nil {
		panic(err)
	}
	contactView, err = template.ParseFiles("contact.gohtml", "layout/footer.gohtml")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	// convert the notfound to http.handler
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}

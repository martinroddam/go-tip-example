package main

import (
	"html/template"
	"net/http"

	"github.com/martinroddam/go-tip"
)

type PageData struct {
	PageTitle string
	PageMsg   string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	data := PageData{
		PageTitle: "Home",
		PageMsg:   "You're back home!",
	}
	tmpl.Execute(w, data)
	gotip.Verify("Navigated to Home")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	data := PageData{
		PageTitle: "Products",
		PageMsg:   "Take a look at our products!",
	}
	tmpl.Execute(w, data)
	gotip.Verify("Navigated to Products")
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/products", productsHandler)

	http.ListenAndServe(":8081", nil)
}

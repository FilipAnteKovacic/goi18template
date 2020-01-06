package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Contain all translations
var i18Map map[string]string

var router *httprouter.Router

// Use to define template functions
var templateFuncs map[string]interface{}

func init() {

	var err error

	// Parse translations
	i18Map, err = parseJSONi18File("en.json")
	if err != nil {
		fmt.Println("error while parsing file", err)
	}

	// Define template funcs
	templateFuncs = template.FuncMap{
		"i18": i18,
	}

}

func main() {

	router := httprouter.New()
	router.GET("/", index)
	router.GET("/login", login)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Init base temaplate
	baseTemplate, err := template.
		New("index.html").
		Funcs(templateFuncs).
		ParseFiles(
			"templates/index.html",
		)
	if err != nil {
		fmt.Println("Error occurred while parsing baseTemplate", err)
		return
	}

	// Cloning from base template
	contentTemplate, err := template.
		Must(baseTemplate.Clone()).
		ParseFiles(
			"templates/content.html",
		)
	if err != nil {
		fmt.Println("Error occurred while parsing content", err)
		return
	}

	err = contentTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error occurred while executing the template or writing its output : ", err)
		return
	}

}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Init another template
	loginTemplate, err := template.
		New("login.html").
		Funcs(templateFuncs).
		ParseFiles(
			"templates/login.html",
		)
	if err != nil {
		fmt.Println("Error occurred while parsing loginTemplate", err)
		return
	}

	err = loginTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error occurred while executing the template or writing its output : ", err)
		return
	}

}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var i18Map map[string]string

var router *httprouter.Router

var templateFuncs map[string]interface{}

func init() {

	var err error
	i18Map, err = parseJSONi18File("en.json")
	if err != nil {
		fmt.Println("error while parsing file", err)
	}

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

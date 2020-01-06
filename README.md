# Go i18 template functions example

Example how to use func in templates
- using simple i18 func 

## Steps

### Router

First step is to init web server for our routes
- in this example github.com/julienschmidt/httprouter is used for router

```
var router *httprouter.Router

func main() {

    router := httprouter.New()
    router.GET("/", index)
    router.GET("/login", login)

    err := http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatal(err)
    }

}
```

- empty route handling funcs for index & login

```
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
```

### i18 functions

Next is adding functions that will going to use in template
- see i18.go file

Add functions to template.FuncMap 

```
var templateFuncs map[string]interface{}

templateFuncs = template.FuncMap{
    "i18": i18,
}
```

### Translation map

Add JSON translations file & parse file in map
- see en.json file

Parsing file in map in init func before app starts

```
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
```

### Template init

Add new folder with templates & add index.html, login.html 
- see templates folder

Parse files in index&login routes function

- index
```
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

	err = baseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error occurred while executing the template or writing its output : ", err)
		return
	}

}
```

- login
```
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
```

### Template cloning

Add antoher template that close baseTemplate

```
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
```

### Using functions in template

In template files i18 functions that is defined is using with given key from json file to have translations

```
<title>{{i18 "login_title"}}</title>

<title>{{i18 "login_title"}}</title>

```

### Next steps

Create new JSON lang file & parse it in init for another language

```
i18Map, err = parseJSONi18File("fr.json")
if err != nil {
    fmt.Println("error while parsing file", err)
}
```
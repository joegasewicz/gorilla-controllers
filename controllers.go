package gorillacontrollers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
)

type CurrentHandler func(http.ResponseWriter, *http.Request, *map[string]interface{})

type GorillaControllers struct {
	CurrentRoute     string
	CurrentHandler   CurrentHandler
	CurrentMethods   []string
	CurrentTemplates []string
	BaseTemplates    []string
	// Generic router e.g Gorilla's *mux.Router
	r            *mux.Router
	TemplateName string
}

func handleFuncWrapper(g *GorillaControllers, t *GTemplate, h CurrentHandler) http.HandlerFunc {
	var data map[string]interface{}

	return func(w http.ResponseWriter, r *http.Request) {
		// Create CurrentTemplates here if exist
		h(w, r, &data)
		templates := t.GTemplates(g.CurrentTemplates...)
		te, err := template.ParseFiles(templates...)
		if err != nil {
			log.Fatalf(err.Error())
		}
		te.ExecuteTemplate(w, "layout", data) // TODO "layout" make dynamic
	}
}

func create(g *GorillaControllers) {
	t := GTemplate{
		BaseTemplates: g.BaseTemplates,
	}
	handler := g.CurrentHandler
	g.r.HandleFunc(g.CurrentRoute, handleFuncWrapper(g, &t, handler)).Methods(g.CurrentMethods...)
}

// New Returns a pointer to a GorillaControllers struct.
// r - pointer to Gorilla's mux.Router
// baseTemplates - is a list of relative paths to your partial templates (not your main route templates)
// templateName - is the name of the base template file e.g 'layout' for 'layout.html'
//
//
//		baseTemplates := []string{
//			"./templates/layout.html",
//			"./templates/sidebar.html",
//			"./templates/navbar.html",
//			"./templates/footer.html",
//		}
//		r := mux.NewRouter()
//		g := gorillacontrollers.New(r, baseTemplates, "layout")
//
func New(r *mux.Router, baseTemplates []string, templateName string) *GorillaControllers {
	return &GorillaControllers{
		CurrentRoute:   "",
		CurrentHandler: nil,
		CurrentMethods: nil,
		BaseTemplates:  baseTemplates,
		r:              r,
		TemplateName:   templateName,
	}
}

// Route A string representing the incoming request URL.
// This is the first argument to Gorilla's mux.Route() method or the first
// argument to http.HandleFunc(). For Example:
//
//
//		g.Route("/") // ... other chained methods
//
func (g *GorillaControllers) Route(route string) *GorillaControllers {
	g.CurrentRoute = route
	return g
}

// Controller is called if the Route request URL is matched.
// handler arg is your controller function.Create a controller - template data needs to be passed
// by value to `data *map[string]interface{}`
//
//
//		func Home(w http.ResponseWriter, r *http.Request, data *map[string]interface{}) {
//    		var templateData map[string]interface{} // Create a map to store your template data
//    		templateData = make(map[string]interface{})
//    		templateData["heading"] = "Create a new advert"
//    		*data = templateData // pass by value back to `data`
//		}
//
//		g.
// 		  // ...
//		  .Controller(Home)
//		  // ...
//
func (g *GorillaControllers) Controller(handler CurrentHandler) *GorillaControllers {
	g.CurrentHandler = handler
	return g
}

// Methods CRUD methods to match on the request URL. For Example
//
// 		g.Methods("GET")
//		g.Methods("GET", "POST", "DELETE")
//
// Example above is just for demonstration, you must call the Route, Controller,
// Template or Init methods.
func (g *GorillaControllers) Methods(methods ...string) *GorillaControllers {
	g.CurrentMethods = methods
	return g
}

// Templates Method that takes a single template relative path or multiple template
// path slices of main route templates (not partial templates). For example:
//
//		 g.Route("/")
//			.controller(Home)
//			.Methods("GET")
//			.Templates("./templates/hero.html", "./templates/routes/home.html")
//
// The above example adds a `hero.html` partial template & a main route `home.html` template.
func (g *GorillaControllers) Templates(templates ...string) *GorillaControllers {
	g.CurrentTemplates = templates
	create(g)
	return g
}

// Init method is used to initiate the chained method calls if you do not register
// the g.Template method. For example:
//
//
//	g.Route("/")
//		.Controller(routes.Home)
//		.Methods("GET")
//		.Init()
//
// In the above example we didn't call the Templates() struct method, therefore Init()
// must be called at the end of the method chain.
func (g *GorillaControllers) Init() *GorillaControllers {
	create(g)
	return g
}

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

func handleFuncWrapper(g *GorillaControllers, t *GTemplate) http.HandlerFunc {
	var data map[string]interface{}
	return func(w http.ResponseWriter, r *http.Request) {
		// Create CurrentTemplates here if exist
		g.CurrentHandler(w, r, &data)
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
	g.r.HandleFunc(g.CurrentRoute, handleFuncWrapper(g, &t)).Methods(g.CurrentMethods...)
	//	g.r.HandleFunc(g.CurrentRoute, handleFuncWrapper(g.CurrentHandler, &t, g.CurrentTemplates)).Methods(g.CurrentMethods...)
}

// New Returns a pointer to a GorillaControllers struct.
// r - pointer to Gorilla's mux.Router
// baseTemplates - is a list of relative paths to your templates
// templateName - is the name of the base template file e.g 'layout' for 'layout.html'
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

func (g *GorillaControllers) Route(route string) *GorillaControllers {
	g.CurrentRoute = route
	return g
}

func (g *GorillaControllers) Controller(handler func(http.ResponseWriter, *http.Request, *map[string]interface{})) *GorillaControllers {
	g.CurrentHandler = handler
	return g
}

func (g *GorillaControllers) Methods(methods ...string) *GorillaControllers {
	g.CurrentMethods = methods
	return g
}

func (g *GorillaControllers) Templates(templates ...string) *GorillaControllers {
	g.CurrentTemplates = templates
	create(g)
	return g
}

func (g *GorillaControllers) Init() *GorillaControllers {
	create(g)
	return g
}

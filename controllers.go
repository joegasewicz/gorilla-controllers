package gorillacontrollers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
)

type GorillaControllers struct {
	CurrentRoute     string
	CurrentHandler   func(http.ResponseWriter, *http.Request, *map[string]interface{})
	CurrentMethods   []string
	CurrentTemplates []string
	BaseTemplates    []string
	// Generic router e.g Gorilla's *mux.Router
	r *mux.Router
}

func handleFuncWrapper(h func(http.ResponseWriter, *http.Request, *map[string]interface{}), t *GTemplate, templatePath []string) http.HandlerFunc {
	var data map[string]interface{}
	return func(w http.ResponseWriter, r *http.Request) {
		// Create CurrentTemplates here if exist
		h(w, r, &data)
		templates := t.GTemplates(templatePath...)
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
	g.r.HandleFunc(g.CurrentRoute, handleFuncWrapper(g.CurrentHandler, &t, g.CurrentTemplates)).Methods(g.CurrentMethods...)
}

func New(r *mux.Router, baseTemplates []string) *GorillaControllers {
	return &GorillaControllers{
		CurrentRoute:   "",
		CurrentHandler: nil,
		CurrentMethods: nil,
		BaseTemplates:  baseTemplates,
		r:              r,
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

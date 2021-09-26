package gorillacontrollers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
)

type GTemplate struct {
	BaseTemplates []string
}

func (t *GTemplate) GTemplates(routeTemplateName ...string) []string {
	for _, template := range routeTemplateName {
		t.BaseTemplates = append(t.BaseTemplates, template)
	}
	return t.BaseTemplates
}

type GorillaControllers struct {
	route         string
	handler       func(http.ResponseWriter, *http.Request, *map[string]interface{})
	methods       []string
	templates     []string
	BaseTemplates []string
	// Generic router e.g Gorilla's *mux.Router
	r *mux.Router
}

func handleFuncWrapper(h func(http.ResponseWriter, *http.Request, *map[string]interface{}), t *GTemplate, templatePath []string) http.HandlerFunc {
	var data map[string]interface{}
	return func(w http.ResponseWriter, r *http.Request) {
		// Create templates here if exist
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
	g.r.HandleFunc(g.route, handleFuncWrapper(g.handler, &t, g.templates)).Methods(g.methods...)
}

func NewGorillaControllers(r *mux.Router, baseTemplates []string) *GorillaControllers {
	return &GorillaControllers{
		route:         "",
		handler:       nil,
		methods:       nil,
		BaseTemplates: baseTemplates,
		r:             r,
	}
}

func (g *GorillaControllers) Route(route string) *GorillaControllers {
	g.route = route
	return g
}

func (g *GorillaControllers) Controller(handler func(http.ResponseWriter, *http.Request, *map[string]interface{})) *GorillaControllers {
	g.handler = handler
	return g
}

func (g *GorillaControllers) Methods(methods ...string) *GorillaControllers {
	g.methods = methods
	return g
}

func (g *GorillaControllers) Templates(templates ...string) *GorillaControllers {
	g.templates = templates
	create(g)
	return g
}

func (g *GorillaControllers) Init() *GorillaControllers {
	create(g)
	return g
}

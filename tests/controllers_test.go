package tests

import (
	"fmt"
	"github.com/gorilla/mux"
	gorillacontrollers "github.com/joegasewicz/gorilla-controllers"
	"net/http"
	"testing"
)

var r *mux.Router
var baseTemplates []string

func setUp() {
	baseTemplates = []string{
		"./templates/layout.html",
		"./templates/sidebar.html",
		"./templates/navbar.html",
		"./templates/footer.html",
	}
	r = mux.NewRouter()
}

func tearDown() {
	baseTemplates = nil
	r = nil
}

func TestNew(t *testing.T) {
	setUp()
	g := gorillacontrollers.New(r, baseTemplates)
	if len(g.BaseTemplates) != 4 {
		t.Errorf("Expected '%v' to be length '%v'", len(g.BaseTemplates), 4)
	}
	tearDown()
}

func TestRoute(t *testing.T) {
	setUp()
	g := gorillacontrollers.New(r, baseTemplates)
	g.Route("/")
	if g.CurrentRoute != "/" {
		t.Errorf("Expected '%v' route but got '%v'", "/", g.CurrentRoute)
	}
	tearDown()
}

func TestController(t *testing.T) {
	setUp()
	g := gorillacontrollers.New(r, baseTemplates)
	_ = g.Controller(Basic)
	chType := fmt.Sprintf("%T", g.CurrentHandler)
	if chType != "func(http.ResponseWriter, *http.Request, *map[string]interface {})" {
		t.Errorf("Expected g.CurrentHandler to be a function not %T", g.CurrentHandler)
	}
	tearDown()
}

func TestMethods(t *testing.T) {
	setUp()
	g := gorillacontrollers.New(r, baseTemplates)
	_ = g.Methods("GET")
	if g.CurrentMethods[0] != "GET" {
		t.Errorf("Expected '%v' bit got '%v", "GET", g.CurrentMethods[0])
	}
	_ = g.Methods("GET", "POST", "DELETE", "PUT")
	if g.CurrentMethods[1] != "POST" &&
		g.CurrentMethods[2] != "DELETE" &&
		g.CurrentMethods[3] != "PUT" {
		t.Errorf("Expected values for CurrentMethods do not match")
	}
	tearDown()
}

func TestTemplates(t *testing.T) {
	setUp()
	g := gorillacontrollers.New(r, baseTemplates)
	_ = g.Templates("index.html")
	if g.CurrentTemplates[0] != "index.html" {
		t.Errorf("g.CurrentTemplates does not contain the correct value")
	}
	tearDown()
}

func TestInit(t *testing.T) {
	setUp()
	g := gorillacontrollers.New(r, baseTemplates)
	g.Route("/").Controller(Basic).Methods("GET", "POST").Init()
	if g.CurrentRoute != "/" && g.CurrentMethods[1] != "POST" {
		t.Errorf("g.Init dit not add correct values to GorillaControllers struct")
	}
	tearDown()
}

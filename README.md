# Gorilla Controllers
Use controllers in the Gorilla Mux library

### Basic Usage

Create a controller
```go
// Create a controller - template data needs to be passed by value to `data *map[string]interface{}`
func Home(w http.ResponseWriter, r *http.Request, data *map[string]interface{}) {
    var templateData map[string]interface{} // Create a map to store your template data 
    templateData = make(map[string]interface{})
    templateData["heading"] = "Create a new advert"
    *data = templateData // pass by value back to `data`
}
```
Library Setup
```go
baseTemplates := []string{
    "./templates/layout.html",
    "./templates/sidebar.html",
    "./templates/navbar.html",
    "./templates/footer.html",
}

r := mux.NewRouter()
g := helpers.NewGorillaControllers(r, baseTemplates)


g.Route("/")
    .Controller(Home)
    .Methods("GET", "POST")
    .Templates("home.html") // If you do not call Templates() then you must call Init() instead

```
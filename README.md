# Gorilla Controllers
Use controllers with the Gorilla Mux library

#### Why controllers?
Gorilla's mux library is a brilliant fully featured mux tool, Gorilla Controllers is a library that replaces the HandleFunc
with a `Controller` function & a `Templates` function. This way you only need to manage your data inside your controller, 
all your template setup logic is now handled by Gorilla Controllers.

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
g := gorillacontrollers.New(r, baseTemplates, "layout") // "layout" is your base template


g.Route("/")
    .Controller(Home) // Controller now replaces Gorilla's HandleFunc
    .Methods("GET", "POST")
    .Templates("./templates/navbar.html", "./templates/home.html") // If you do not call Templates() then you must call Init() instead

```

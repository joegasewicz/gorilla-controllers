package tests

import "net/http"

func Basic(w http.ResponseWriter, r *http.Request, data *map[string]interface{}) {
	var templateData map[string]interface{} // Create a map to store your template data
	templateData = make(map[string]interface{})
	templateData["heading"] = "Create a new advert"
	*data = templateData // pass by value back to `data`
}

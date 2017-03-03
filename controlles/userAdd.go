package controlles

import (
	"MongoDB_CURD/models"
	"encoding/json"
	"fmt"
	"net/http"
)

var mapResult = make(map[string]interface{})
var mapRequest = make(map[string]interface{})
var result bool
var info string

func AddUser(w http.ResponseWriter, r *http.Request) {
	result = true
	params := r.URL.Query()
	requestJson := params.Get("json")
	if len(requestJson) == 0 {
		result, info = models.AddUser(
			params.Get("name"),
			params.Get("password"),
			params.Get("age"),
			params.Get("phone"),
		)
	} else {
		err := json.Unmarshal([]byte(requestJson), &mapRequest)
		if err != nil {
			result = false
		} else {
			result, info = models.AddUser(
				mapRequest["name"].(string),
				mapRequest["password"].(string),
				mapRequest["age"].(string),
				mapRequest["phone"].(string),
			)
		}

	}
	mapResult["status"] = result
	mapResult["info"] = info
	json, _ := json.Marshal(mapResult)
	fmt.Fprint(w, string(json))

}

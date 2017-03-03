package controlles

import (
	"net/http"
	"MongoDB_CURD/models"
	"encoding/json"
	"fmt"
)

func UserLognin(w http.ResponseWriter, r *http.Request) {
	result = true
	params := r.URL.Query()
	requestJson := params.Get("json")
	if len(requestJson) == 0 {
		result, info = models.UserLogin(
			params.Get("name"),
			params.Get("password"),
		)
	} else {
		err := json.Unmarshal([]byte(requestJson), &mapRequest)
		if err != nil {
			result = false
		} else {
			result, info = models.UserLogin(
				mapRequest["name"].(string),
				mapRequest["password"].(string),
			)
		}

	}
	mapResult["status"] = result
	mapResult["info"] = info
	json, _ := json.Marshal(mapResult)
	fmt.Fprint(w, string(json))

}

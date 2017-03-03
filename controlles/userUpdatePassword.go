package controlles

import (
	"MongoDB_CURD/models"
	"encoding/json"
	"fmt"
	"net/http"
)



func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	result = true
	params := r.URL.Query()
	requestJson := params.Get("json")
	if len(requestJson) == 0 {
		result,info = models.UpdateUserPassword(
			params.Get("name"),
			params.Get("password"),
		)
	} else {
		err := json.Unmarshal([]byte(requestJson), &mapRequest)
		if err != nil {
			result = false
		} else {
			result,info = models.UpdateUserPassword(
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


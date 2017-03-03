package controlles

import (
	"MongoDB_CURD/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func DeleteUserByName(w http.ResponseWriter, r *http.Request) {
	result = true
	params := r.URL.Query()
	requestJson := params.Get("json")
	if len(requestJson) == 0 {
		result,info = models.DeleteUserByName(
			params.Get("name"),
		)
	} else {
		err := json.Unmarshal([]byte(requestJson), &mapRequest)
		if err != nil {
			result = false
		} else {
			result,info = models.DeleteUserByName(
				mapRequest["name"].(string),
			)
		}

	}
	mapResult["status"] = result
	mapResult["info"] = info
	json, _ := json.Marshal(mapResult)
	fmt.Fprint(w, string(json))

}

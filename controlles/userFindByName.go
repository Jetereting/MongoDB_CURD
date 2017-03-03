package controlles

import (
	"MongoDB_CURD/models"
	"encoding/json"
	"fmt"
	"net/http"
)



func FindUserByName(w http.ResponseWriter, r *http.Request) {
	users:=[]models.User{}
	var err error
	result=true
	params := r.URL.Query()
	requestJson := params.Get("json")
	if len(requestJson) == 0 {
		users,err= models.FindUserByName(
			params.Get("name"),
		)
		if err != nil {
			result = false
		}
	} else {
		err = json.Unmarshal([]byte(requestJson), &mapRequest)
		if err != nil {
			result = false
		} else {
			users,err = models.FindUserByName(
				mapRequest["name"].(string),
			)
		}

	}
	if result&&users!=nil{
		json, _ := json.Marshal(users)
		fmt.Fprint(w, string(json))
	}else {
		mapResult["status"] = false
		mapResult["info"] = "没有该用户"
		json, _ := json.Marshal(mapResult)
		fmt.Fprint(w, string(json))
	}
}


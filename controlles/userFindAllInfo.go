package controlles

import (
	"MongoDB_CURD/models"
	"encoding/json"
	"fmt"
	"net/http"
)



func FindAllUserInfo(w http.ResponseWriter, r *http.Request) {
	result,err:= models.FindAllUserInfo()
	if err!=nil{
		mapResult["status"] = false
		mapResult["info"] = "没有一个用户"
		json, _ := json.Marshal(mapResult)
		fmt.Fprint(w, string(json))
	}else {
		json, _ := json.Marshal(result)
		fmt.Fprint(w, string(json))
	}

}


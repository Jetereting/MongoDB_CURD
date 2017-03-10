package controlles

import (
	"net/http"
	"encoding/json"
	"fmt"
	"MongoDB_CURD/models"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	mapRequest := make(map[string]interface{})

	err:= json.NewDecoder(r.Body).Decode(&mapRequest)
	if err!=nil{
		fmt.Println("Add User Error:",err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, info:= models.AddUser(
		mapRequest["name"].(string),
		mapRequest["password"].(string),
		mapRequest["age"].(string),
		mapRequest["phone"].(string),
	)

	mapResult:=map[string]interface{}{
		"status":result,
		"info":info,
	}

	json.NewEncoder(w).Encode(mapResult)
}
func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	name:=r.URL.Query().Get("name")
	result, info:=models.DeleteUserByName(name)

	mapResult:=map[string]interface{}{
		"status":result,
		"info":info,
	}

	json.NewEncoder(w).Encode(mapResult)
}

func FindAllUser(w http.ResponseWriter, r *http.Request)  {
	users,err:=models.FindAllUserInfo()
	if(err!=nil){
		fmt.Println("FindAllUser ERROR:",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func FindUserByName(w http.ResponseWriter, r *http.Request)  {
	name:=r.URL.Query().Get("name")
	users,err:=models.FindUserByName(name)
	if(err!=nil){
		fmt.Println("FindAllUser ERROR:",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func Login(w http.ResponseWriter, r *http.Request)  {
	mapRequest := make(map[string]interface{})
	err:= json.NewDecoder(r.Body).Decode(&mapRequest)
	if err!=nil{
		fmt.Println("Login Error:",err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, info:= models.UserLogin(
		mapRequest["name"].(string),
		mapRequest["password"].(string),
	)

	mapResult:=map[string]interface{}{
		"status":result,
		"info":info,
	}

	json.NewEncoder(w).Encode(mapResult)
}

func UpdateUserPassword(w http.ResponseWriter,r *http.Request)  {
	mapRequest := make(map[string]interface{})
	err:= json.NewDecoder(r.Body).Decode(&mapRequest)
	if err!=nil{
		fmt.Println("Login Error:",err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, info:= models.UpdateUserPassword(
		mapRequest["name"].(string),
		mapRequest["password"].(string),
	)

	mapResult:=map[string]interface{}{
		"status":result,
		"info":info,
	}

	json.NewEncoder(w).Encode(mapResult)
}


package routers

import (
	"net/http"
	"MongoDB_CURD/controlles"
	"MongoDB_CURD/conf"
	"fmt"
	"github.com/drone/routes"
)

func init()  {

	mux:=routes.New()
	mux.Get("/",controlles.Index)
	mux.Get("/user/findAllUserInfo",controlles.FindAllUser)
	mux.Get("/user/findUserByName",controlles.FindUserByName)
	mux.Post("/user/updateUserPassword",controlles.UpdateUserPassword)
	mux.Del("/user/deleteUserByName",controlles.DeleteUser)
	mux.Put("/user/addUser",controlles.AddUser)
	mux.Get("/user/userLogin",controlles.Login)
	http.Handle("/",mux)

	//http.HandleFunc("/", controlles.Index)
	//http.HandleFunc("/addUser", controlles.AddUser)
	//http.HandleFunc("/deleteUserByName", controlles.DeleteUserByName)
	//http.HandleFunc("/findAllUserInfo", controlles.FindAllUserInfo)
	//http.HandleFunc("/findUserByName", controlles.FindUserByName)
	//http.HandleFunc("/updateUserPassword", controlles.UpdateUserPassword)

	err:=http.ListenAndServe(conf.Port,nil)
	if err!=nil{
		fmt.Println("ERROR:",err)
		return
	}
}
package models

import (
	"MongoDB_CURD/conf"
	"MongoDB_CURD/utils"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
)

var session *mgo.Session

type User struct {
	ID string `bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
	Age      string `json:"age"`
	Phone    string `json:"phone"`
}

func init() {

	fmt.Println("server start!!")
	var err error
	session, err = mgo.Dial(conf.DbUrl)
	if err != nil {
		fmt.Println("Dial fail!")
		return
	}
	session.SetMode(mgo.Monotonic, true)
}

func Destroy() {
	session.Close()
}

func cHandle(handle func(c *mgo.Collection)) {
	ns := session.Clone()
	defer ns.Close()
	handle(ns.DB(conf.DbName).C("user"))
}

func AddUser(name, password, age, phone string) (bool, string) {
	result:=true
	info:="插入成功"
	cHandle(func(c *mgo.Collection) {
		n,err:= c.Find(bson.M{"name": name}).Count()
		if err!=nil{
			fmt.Println("model-AddUser ERROR:",err)
			result = false
			info = "插入错误请联系管理员"
			return
		}else if n>0{
			result = false
			info = "已存在该用户名"
			return
		}

		user:=&User{
			bson.NewObjectId().Hex(),
			name,
			utils.DoubleMd5(password),
			age,
			phone,
		}

		err = c.Insert(user)
		if err != nil {
			fmt.Println("model-AddUser ERROR:",err)
			result = false
			info = "插入失败"
		}

	})
	return result, info
}

func UserLogin(name, password string) (bool, string) {
	result:=true
	info:="登陆成功"
	cHandle(func(c *mgo.Collection) {
		user:=User{}
		err:=c.Find(bson.M{"name":name}).One(user)
		if err!=nil{
			fmt.Println("model-UserLogin ERROR:",err)
			result=false
			info="没有该用户"
			return
		}
		if strings.Compare(utils.DoubleMd5(password),user.Password)!=0{
			result=false
			info="密码错误"
		}

	})
	return result, info
}

func DeleteUserByName(name string) (bool, string){
	info:="删除成功"
	result := true
	cHandle(func(c *mgo.Collection) {
		_, err := c.RemoveAll(bson.M{"name": name})
		if err != nil {
			fmt.Println("model-DeleteUserByName",err)
			result = false
			info="删除失败"
			return
		}
	})
	return result,info
}

func UpdateUserPassword(name, password string) (bool, string) {
	info:="更新成功"
	result := true
	cHandle(func(c *mgo.Collection) {
		err := c.Update(bson.M{"name": name}, bson.M{"$set": bson.M{"password": password}})
		if err != nil {
			fmt.Println("model-UpdateUserPassword",err)
			result = false
			info="更新失败"
			return
		}
	})
	return result,info
}

func FindAllUserInfo() (users []User, err error) {
	cHandle(func(c *mgo.Collection) {
		err = c.Find(nil).All(&users)
		if err!=nil{
			fmt.Println("model-FindAllUserInfo",err)
			return
		}
	})
	return users, err
}

func FindUserByName(name string) (users []User, err error) {
	cHandle(func(c *mgo.Collection) {
		err = c.Find(bson.M{"name": name}).All(&users)
		if err!=nil{
			fmt.Println("model-FindAllUserInfo",err)
			return
		}
	})
	return users, err
}

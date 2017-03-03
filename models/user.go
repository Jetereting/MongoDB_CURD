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
var err error
var result = true
var info = ""
var user User

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Age      string `json:"age"`
	Phone    string `json:"phone"`
}

func init() {

	fmt.Println("server start!!")

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
	cHandle(func(c *mgo.Collection) {
		if c.Find(bson.M{"name": name}).One(&User{}) == nil {
			result = false
			info = "该用户名已存在"
		} else {
			err := c.Insert(&User{name, utils.DoubleMd5(password), age, phone})
			if err != nil {
				result = false
				info = "插入失败"
			}
		}

	})
	return result, info
}

func UserLogin(name, password string) (bool, string) {
	info="登陆成功"
	cHandle(func(c *mgo.Collection) {
		if c.Find(bson.M{"name": name}).One(&user)!= nil {
			result = false
			info = "没有该用户"
		} else {
			if strings.EqualFold(utils.DoubleMd5(password),user.Password){
				result=true
			}else {
				result = false
				info = "密码错误"
			}
		}

	})
	return result, info
}

func DeleteUserByName(name string) (bool, string){
	info="删除成功"
	result := true
	cHandle(func(c *mgo.Collection) {
		_, err := c.RemoveAll(bson.M{"name": name})
		if err != nil {
			result = false
			info="删除失败"
		}
	})
	return result,info
}

func UpdateUserPassword(name, password string) (bool, string) {
	info="更新成功"
	result := true
	cHandle(func(c *mgo.Collection) {
		err := c.Update(bson.M{"name": name}, bson.M{"$set": bson.M{"password": password}})
		if err != nil {
			result = false
			info="更新失败"
		}
	})
	return result,info
}

func FindAllUserInfo() (users []User, err error) {
	cHandle(func(c *mgo.Collection) {
		err = c.Find(nil).All(&users)
	})
	return users, err
}

func FindUserByName(name string) (users []User, err error) {
	cHandle(func(c *mgo.Collection) {
		err = c.Find(bson.M{"name": name}).All(&users)
	})
	return users, err
}

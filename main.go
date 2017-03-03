package main

import (
	"MongoDB_CURD/models"
	_"MongoDB_CURD/routers"
)


func main() {
	defer models.Destroy()
}

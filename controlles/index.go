package controlles

import (
	"net/http"
	"fmt"
)


func Index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"Index Page,welcome.")
}

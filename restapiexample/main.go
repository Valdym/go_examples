package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/Valdym/go_examples/restapiexample/models"
	"github.com/Valdym/go_examples/restapiexample/utils"
)

func main() {
	handler := &utils.RegexpHandler{}

	//Welcome message get!
	handler.HandleFunc(regexp.MustCompile("^/api/$"), models.HomeAPI)

	handler.HandleFunc(regexp.MustCompile("^/api/users/$"), models.Employee)
	fmt.Println(regexp.QuoteMeta(`/api/users/\d+$`))
	fmt.Println("/api/users/" + "^[0-9]+$")
	re := regexp.MustCompile("^/api/users/[1-9]+$")
	handler.HandleFunc(re, models.PutEmployee)

	//If above endpoints fail this should work!
	//https://pkg.go.dev/net/http#ServeMux
	//"Note that since a pattern ending in a slash names a rooted subtree,
	//the pattern "/" matches all paths not matched by other registered patterns."
	handler.HandleFunc(regexp.MustCompile("/"), func(w http.ResponseWriter, r *http.Request) {
		resp := utils.Response{Resp: w}
		resp.Text(http.StatusNotFound, "Not found", "text/plain")
	})
	http.ListenAndServe(":8080", handler)
}

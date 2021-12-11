package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/Valdym/go_examples/restapiexample/models"
	"github.com/Valdym/go_examples/restapiexample/utils"
)

func main() {
	handler := &utils.RegexpHandler{}

	//Welcome message get!
	handler.HandleFunc(regexp.MustCompile("^/api/$"), homeAPI)

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

func homeAPI(w http.ResponseWriter, r *http.Request) {
	message1 := API{WelcomeMessage: "API Home"}
	body, err := json.Marshal(message1)

	utils.CheckError(err, w, r)
	fmt.Println(string(body))

	resp := utils.Response{Resp: w}
	resp.Text(http.StatusOK, string(body), "text/json")
}

type API struct {
	WelcomeMessage string `json:message`
}

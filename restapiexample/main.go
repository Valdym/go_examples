package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Valdym/go_examples/restapiexample/models"
	"github.com/Valdym/go_examples/restapiexample/utils"
)

func main() {

	apiRoot := "/api"

	//Welcome message get!
	http.HandleFunc(apiRoot, homeAPI)

	http.HandleFunc(apiRoot+"/users/", models.Employee)

	err := http.ListenAndServe(":8080", nil)
	utils.CheckError(err)
}

func homeAPI(w http.ResponseWriter, r *http.Request) {
	message1 := API{WelcomeMessage: "API Home"}
	body, err := json.Marshal(message1)

	utils.CheckError(err)
	fmt.Println(string(body))

	resp := utils.Response{Resp: w}
	resp.Text(http.StatusOK, string(body), "text/json")
}

type API struct {
	WelcomeMessage string `json:message`
}

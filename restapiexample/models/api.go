package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Valdym/go_examples/restapiexample/utils"
)

func HomeAPI(w http.ResponseWriter, r *http.Request) {
	message1 := API{WelcomeMessage: "API Home"}
	body, err := json.Marshal(message1)

	utils.CheckError(err, w, r)
	fmt.Println(string(body))

	resp := utils.Response{Resp: w}
	resp.Text(http.StatusOK, string(body), "text/json")
}

type API struct {
	WelcomeMessage string `json:"message"`
}

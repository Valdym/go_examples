package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Valdym/go_examples/restapiexample/utils"
)

var list_employee []Employees

type Employees struct {
	Id         int    `json:id`
	FirstName  string `json:firstname`
	LastName   string `json:lastname`
	Department string `json:department`
	Age        int    `json:age`
	Salary     int    `json:salary`
}

func Employee(w http.ResponseWriter, r *http.Request) {
	var e Employees
	defer r.Body.Close()

	switch r.Method {
	case http.MethodGet:
		resp := utils.Response{Resp: w}
		body, err := json.Marshal(list_employee)
		utils.CheckError(err)
		resp.Text(http.StatusOK, string(body), "text/plain")
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&e)
		utils.CheckError(err)
		list_employee = append(list_employee, e)
		fmt.Println(list_employee)
		resp := utils.Response{Resp: w}
		resp.Text(http.StatusOK, "Operation Successful! \n"+e.FirstName+" "+e.LastName+" is added to our Company!", "text/plain")
	default:
		resp := utils.Response{Resp: w}
		resp.Text(http.StatusMethodNotAllowed, "Status not allowed", "text/plain")
	}

	//idvalue := r.Body
}

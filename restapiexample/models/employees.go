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
	Firstname  string `json:first_name`
	LastName   string `json:last_name`
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
		resp.Text(http.StatusOK, "Employee Info:\n", "text/plain")
	case http.MethodPost:
		err := json.NewDecoder(r.Body).Decode(&e)
		utils.CheckError(err)
		fmt.Println(e.Firstname)
		list_employee = append(list_employee, e)
		fmt.Println(list_employee)
		resp := utils.Response{Resp: w}
		resp.Text(http.StatusOK, "Operation Successful! \n"+e.Firstname+" is added to our Company!", "text/plain")
	default:
		resp := utils.Response{Resp: w}
		resp.Text(http.StatusMethodNotAllowed, "Status not allowed", "text/plain")
	}

	//idvalue := r.Body
}

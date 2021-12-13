package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Valdym/go_examples/restapiexample/utils"
)

var list_employee []Employees

type Employees struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Department string `json:"department"`
	Age        int    `json:"age"`
	Salary     int    `json:"salary"`
}

func PutEmployee(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	switch r.Method {
	case http.MethodPut:
		var e Employees
		err := json.NewDecoder(r.Body).Decode(&e)
		utils.CheckError(err, w, r)

		re := regexp.MustCompile("[0-9]+")
		idvaluestr := re.FindString(r.URL.Path) //get id value
		idvalueint, _ := strconv.Atoi(idvaluestr)
		for idx, elem := range list_employee {
			if elem.Id == idvalueint {
				list_employee[idx].Id = e.Id
				list_employee[idx].FirstName = e.FirstName
				list_employee[idx].LastName = e.LastName
				list_employee[idx].Department = e.Department
				list_employee[idx].Age = e.Age
				list_employee[idx].Salary = e.Salary
				fmt.Println(elem)
				resp := utils.Response{Resp: w}
				body, err := json.Marshal(list_employee[idx])
				utils.CheckError(err, w, r)
				resp.Text(http.StatusOK, string(body), "text/json")
				return
			}
		}
		resp := utils.Response{Resp: w}
		resp.Text(http.StatusNotFound, "There is no record with given id", "text/plain")
	case http.MethodDelete:
		re := regexp.MustCompile("[0-9]+")
		idvaluestr := re.FindString(r.URL.Path) //get id value
		idvalueint, _ := strconv.Atoi(idvaluestr)
		for idx, elem := range list_employee {
			if elem.Id == idvalueint {
				list_employee = append(list_employee[:idx], list_employee[idx+1:]...) //Delete the member, add remaining members to list
				resp := utils.Response{Resp: w}
				body, err := json.Marshal(elem)
				utils.CheckError(err, w, r)
				resp.Text(http.StatusOK, string(body), "text/json")
				return
			}
		}
		resp := utils.Response{Resp: w}
		resp.Text(http.StatusNotFound, "There is no record with given id", "text/plain")
	default:
		resp := utils.Response{Resp: w}
		resp.Text(http.StatusMethodNotAllowed, "Method not allowed", "text/plain")
	}
}

func Employee(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	switch r.Method {
	case http.MethodGet:
		resp := utils.Response{Resp: w}
		body, err := json.Marshal(list_employee)
		utils.CheckError(err, w, r)
		resp.Text(http.StatusOK, string(body), "text/json")
	case http.MethodPost:
		var e Employees
		err := json.NewDecoder(r.Body).Decode(&e)
		utils.CheckError(err, w, r)
		if err == nil {
			for _, elem := range list_employee {
				if elem.Id == e.Id {
					resp := utils.Response{Resp: w}
					resp.Text(http.StatusConflict, "User with the ID:"+strconv.Itoa(e.Id)+" already exists!", "text/plain")
					return
				}
			}
			list_employee = append(list_employee, e)
			resp := utils.Response{Resp: w}
			body, err := json.Marshal(e)
			utils.CheckError(err, w, r)
			resp.Text(http.StatusOK, string(body), "text/json")
		}

	default:
		resp := utils.Response{Resp: w}
		resp.Text(http.StatusMethodNotAllowed, "Method not allowed", "text/plain")
	}
}

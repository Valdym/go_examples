package models

import (
	"net/http"

	"github.com/Valdym/go_examples/restapiexample/utils"
)

type Employees struct {
	Firstname  string
	LastName   string
	Department string
	Age        int
	Salary     int
}

func Employee(w http.ResponseWriter, r *http.Request) {
	resp := utils.Response{Resp: w}
	resp.Text(http.StatusOK, "Employee Info:\n", "text/plain")

	//idvalue := r.Body
}

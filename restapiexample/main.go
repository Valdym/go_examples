package main

import (
	"net/http"

	"github.com/Valdym/go_examples/restapiexample/models"
	"github.com/Valdym/go_examples/restapiexample/utils"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//Welcome message get!
	router.HandleFunc(("/api/"), models.HomeAPI) // Home API

	router.HandleFunc("/api/users/", models.Employee) //Get All Users!

	router.HandleFunc("/api/users/{id:[0-9]+}", models.PutEmployee) //post,get,delete requests will come here

	//If above endpoints fail this should work!
	//https://pkg.go.dev/net/http#ServeMux
	//"Note that since a pattern ending in a slash names a rooted subtree,
	//the pattern "/" matches all paths not matched by other registered patterns."
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp := utils.Response{Resp: w}
		resp.Text(http.StatusNotFound, "Not found", "text/plain")
	})
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

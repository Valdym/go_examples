package utils

import "net/http"

//Panics if error persists also returns Bad Request assuming request format of JSON is corrupted
func CheckError(err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		resp := Response{Resp: w}
		resp.Text(http.StatusBadRequest, "Invalid JSON Format", "text/plain")
		panic(err)
	}
}

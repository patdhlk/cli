package http

import (
	"encoding/json"
	"net/http"
)

func Handler() http.Handler {
	mux := http.NewServeMux()
	//the handle functions goes here
	mux.Handle("/v1/user/", handleUser())

	return mux //mux implements http.Handler interface, so we can use it in the ListenAndServe method
}

func handleUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET": //handle HTTP.GET
			handleUserGet(w, r)
		case "PUT": //handle HTTP.PUT
			handleUserPut(w, r)
		case "POST": //handle HTTP.POST
			handleUserPost(w, r)
		case "DELETE": //handle HTTP.DELETE
			handleUserDelete(w, r)
		default:
			respondError(w, http.StatusMethodNotAllowed, nil)
		}
	})
}

func handleUserGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World")) //this should be json data
}

func handleUserPut(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusForbidden, nil)
}

func handleUserPost(w http.ResponseWriter, r *http.Request) {
	respondOk(w, nil)
}

func handleUserDelete(w http.ResponseWriter, r *http.Request) {

}

func respondError(w http.ResponseWriter, status int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
}

func respondOk(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-Type", "application/json")

	if body == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		enc.Encode(body)
	}
}

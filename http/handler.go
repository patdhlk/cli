package http

import (
	"net/http"
)

func Handler() http.Handler {
	mux := http.NewServeMux()
	//the handle functions goes here
	//mux.Handle("(/v1/user/", handleUser)

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		mux.ServeHTTP(w, req)
		return
	})

	return handler
}

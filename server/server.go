package server

import "net/http"

func Server() {
	err := http.ListenAndServe(":81", nil)
	if err != nil {
		panic(err)
	}
}

package server

import "net/http"

func server() {
	err := http.ListenAndServe(":81", nil)
	if err != nil {
		panic(err)
	}
}

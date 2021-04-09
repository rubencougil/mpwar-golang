package main

import "net/http"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		_, _ = w.Write([]byte("Hello World!\n"))
	case "DELETE":
		_, _ = w.Write([]byte("Goodbye World!\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		_, _ = w.Write([]byte("Method not Allowed\n"))
	}
}

func main() {
	http.HandleFunc("/", helloWorld)
	_ = http.ListenAndServe(":8000", nil)
}

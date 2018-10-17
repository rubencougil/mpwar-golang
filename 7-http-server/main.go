package main

import "net/http"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			w.Write([]byte("Hello World!\n"))
		case "DELETE":
			w.Write([]byte("Goodbye World!\n"))
		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte("Method not Allowed\n"))
	}
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)
}

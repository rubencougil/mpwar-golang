package main

import "net/http"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!\n"))
}

func main() {
	http.HandleFunc("/", helloWorld)
	_ = http.ListenAndServe(":8000", nil)
}

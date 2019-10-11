package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}", ArticlesCategoryHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

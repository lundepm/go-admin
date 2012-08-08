package main

import (
	"net/http"
	"code.google.com/p/gorilla/mux"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/route/", routeHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8088", nil)
}

package main

import "net/http"

func main() {
	session, err := mgo.Dial("mongodb")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	http.ListenAndServe("0.0.0.0:3000", nil)
}

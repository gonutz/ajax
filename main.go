package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadFile("served_files/hello_world.txt")
		w.Write(data)
	})
	http.Handle("/", http.FileServer(http.Dir("served_files")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

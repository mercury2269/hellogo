package main

import "net/http"

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8081", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

type weatherData struct {
	Name string 'json:"name"'
	Main struct {
		Kelvin float64 'json:"temp"'
	} 'json:"main"'
}
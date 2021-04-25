package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
	fmt.Printf("%T\n", h)
	a := r.Header["Accept-Encoding"]
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	b := r.Header.Get("Accept-Encoding")
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	c := r.Header.Values("Accept-Encoding")
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}

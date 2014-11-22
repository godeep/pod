package main

import (
	"net/http"

	"github.com/squiidz/pod"
)

func main() {
	od := pod.NewPod()
	test := pod.NewSchema(Middle, AuthMid)
	auth := pod.NewSchema(AuthMid, Middle)

	http.Handle("/", od.Fuse(Default))
	http.Handle("/home", od.Fuse(Home).Schema(auth, test).Add(Second))

	http.ListenAndServe(":9000", nil)
}

func Default(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Default Handler\n"))
}

func Home(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Home Handler\n"))
}

func Second(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Second Middleware\n"))
}

func AuthMid(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("AuthMid Middleware\n"))
}

func Middle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("- Middle\n"))
		next.ServeHTTP(rw, req)
	})
}

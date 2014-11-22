package main

import (
	"net/http"

	"github.com/squiidz/fur/middle"
	"github.com/squiidz/pod"
)

func main() {
	p := pod.NewPod()

	p.Glob(GlobalMiddle, middle.Logger)

	http.Handle("/home", p.Fuse(HomeHandler).Add(HomeMiddle, Middle))
	http.Handle("/", p.Fuse(Default))

	http.ListenAndServe(":8080", nil)
}

func GlobalMiddle(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- GlobalMiddlware\n"))
}

func Default(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- Default"))
}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- HomeHandler\n"))
}

func HomeMiddle(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- HomeMiddleware\n"))
}

func Middle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("- Middle\n"))
		next.ServeHTTP(rw, req)
	})
}

package main

import (
	"net/http"

	"github.com/squiidz/pod"
)

func main() {
	p := pod.NewPod()

	p.Glob(GlobalMiddle)

	http.Handle("/home", p.Fuse(HomeHandler).Add(HomeMiddle))

	http.ListenAndServe(":8080", nil)
}

func GlobalMiddle(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- GlobalMiddlware\n"))
}

func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- HomeHandler\n"))
}

func HomeMiddle(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- HomeMiddleware\n"))
}

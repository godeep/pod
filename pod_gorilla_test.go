package pod

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func TestCustom(t *testing.T) {
	// Create a new Pod instance
	p := NewPod()
	// Create a new Gorilla.Mux instance
	m := mux.NewRouter()

	// Add some global middleware
	p.Glob(test3Middle)

	// Merge the Handler with the Global middleware and add a one for this handler
	m.Handle("/", p.Fuse(test2Handler).Add(test4Middle))

	// Start Listening
	http.ListenAndServe(":9000", m)
}

func test2Handler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- Handler (Custom)\n"))
}

func test3Middle(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- Middleware3 (Custom)\n"))
}

func test4Middle(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- Middleware4 (Custom)\n"))
}

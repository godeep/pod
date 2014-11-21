package pod

import (
	"net/http"
	"testing"
)

func TestPod(t *testing.T) {
	// Create a Pod Instance
	p := NewPod()
	// Add Global MiddleWare to your Pod
	p.Glob(testWare, test2Ware)

	// Use Glob Middleware and add some specific one to this handler
	http.Handle("/", p.Fuse(testHandler).Add(NewMid))
	// Just use the Global Ones
	http.Handle("/a", p.Fuse(testHandler))
	// No MiddleWare at All
	http.Handle("/b", http.HandlerFunc(testHandler))

	// Start Listening
	http.ListenAndServe(":9000", nil)

}

func testHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- Handler\n"))
}

func testWare(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- MiddleWare\n"))
}

func test2Ware(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- MiddleWare2\n"))
}

func NewMid(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("- Not Global Middleware\n"))
}

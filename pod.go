package pod

import (
	"fmt"
	"net/http"
	"reflect"
)

type Handler func(rw http.ResponseWriter, req *http.Request)

// PodFunc redefine http.HandlerFunc
type PodFunc http.HandlerFunc

func (p PodFunc) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	p(rw, req)
}

// Middleware is the signature of a valid middleware with Pod
type MiddleWare func(http.Handler) http.Handler

// Pod is the array of the Global Middleware
type Pod []MiddleWare

// NewPod create a new empty Pod
func NewPod() *Pod {
	return &Pod{}
}

// Glob add some Global middleware to the Pod array
func (p *Pod) Glob(m ...interface{}) {
	if len(m) > 0 {
		for _, f := range m {
			switch v := f.(type) {
			case func(http.ResponseWriter, *http.Request):
				*p = append(*p, mutate(http.HandlerFunc(v)))
			case func(http.Handler) http.Handler:
				*p = append(*p, v)
			default:
				fmt.Println("[x] [", reflect.TypeOf(v), "] is not a valid MiddleWare Type.")
			}
		}
	}
}

// Fuse, merge all the global middleware with the provided http.HandlerFunc
func (p *Pod) Fuse(h http.HandlerFunc) PodFunc {
	if len(*p) > 0 {
		var stack http.Handler
		for i, m := range *p {
			switch i {
			case 0:
				stack = m(h)
			default:
				stack = m(stack)
			}
		}

		return PodFunc(stack.(http.HandlerFunc))
	}
	return PodFunc(h)
}

// Add some middleware to a particular handler
func (p PodFunc) Add(m ...http.HandlerFunc) http.Handler {
	var n http.Handler
	if m != nil {
		for _, x := range m {
			mi := mutate(x)
			n = mi(PodFunc(p))
		}
	}
	return n
}

// Mutate generate a valid handler with a provided http.HandlerFunc
func mutate(h http.HandlerFunc) MiddleWare {
	return func(next http.Handler) http.Handler {
		return PodFunc(func(rw http.ResponseWriter, req *http.Request) {
			h(rw, req)
			next.ServeHTTP(rw, req)
		})
	}
}

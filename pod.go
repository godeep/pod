package pod

import (
	"net/http"
)

// PodHandler redefine http.Handler
type PodHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// PodFunc redefine http.HandlerFunc
type PodFunc http.HandlerFunc

func (h PodFunc) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h(rw, req)
}

// Middleware is the signature of a valid middleware with Pod
type MiddleWare func(PodHandler) PodHandler

// Pod is the array of the Global Middleware
type Pod []MiddleWare

// NewPod create a new empty Pod
func NewPod() *Pod {
	return &Pod{}
}

// Glob add some Global middleware to the Pod array
func (p *Pod) Glob(m ...http.HandlerFunc) {
	if m != nil {
		for _, f := range m {
			*p = append(*p, mutate(f))
		}
	}
}

// Fuse, merge all the global middleware with the provided http.HandlerFunc
func (p *Pod) Fuse(h http.HandlerFunc) PodFunc {
	var stack PodHandler
	for i, m := range *p {
		switch i {
		case 0:
			stack = m(h)
		default:
			stack = m(stack)
		}
	}

	return stack.(PodFunc)
}

// Add some middleware to a particular handler
func (h PodFunc) Add(m ...http.HandlerFunc) http.Handler {
	var n http.Handler
	if m != nil {
		for _, x := range m {
			mi := mutate(x)
			n = mi(PodFunc(h))
		}
	}
	return n
}

// Mutate generate a valid handler a provided http.HandlerFunc
func mutate(h http.HandlerFunc) MiddleWare {
	return func(next PodHandler) PodHandler {
		return PodFunc(func(rw http.ResponseWriter, req *http.Request) {
			h(rw, req)
			next.ServeHTTP(rw, req)
		})
	}
}

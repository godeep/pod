package pod

import (
	"net/http"
)

// PodFunc redefine http.HandlerFunc
type PodFunc http.HandlerFunc

func (h PodFunc) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h(rw, req)
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
func (p *Pod) Glob(m ...http.HandlerFunc) {
	if m != nil {
		for _, f := range m {
			*p = append(*p, mutate(f))
		}
	}
}

// Fuse, merge all the global middleware with the provided http.HandlerFunc
func (p *Pod) Fuse(h http.HandlerFunc) PodFunc {
	var stack http.Handler
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
	return func(next http.Handler) http.Handler {
		return PodFunc(func(rw http.ResponseWriter, req *http.Request) {
			h(rw, req)
			next.ServeHTTP(rw, req)
		})
	}
}

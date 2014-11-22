/**********************************
***  Middleware Chaining in Go  ***
***  Code is under MIT license  ***
***    Code by CodingFerret     ***
*** 	github.com/squiidz      ***
***********************************/

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
type Pod struct {
	Handlers []MiddleWare
}

// NewPod create a new empty Pod
func NewPod(m ...interface{}) *Pod {
	p := &Pod{}
	p.wrap(m)
	return p
}

// wrap add some Global middleware to the Pod.Handlers array
func (p *Pod) wrap(m []interface{}) {
	stack := toMiddleware(m)
	for _, s := range stack {
		p.Handlers = append(p.Handlers, s)
	}
}

// Fuse, merge all the global middleware with the provided http.HandlerFunc
func (p *Pod) Fuse(h http.HandlerFunc) *PodFunc {
	if len(p.Handlers) > 0 {
		var stack http.Handler
		for i, m := range p.Handlers {
			switch i {
			case 0:
				stack = m(h)
			default:
				stack = m(stack)
			}
		}
		return stack.(*PodFunc)
	}

	ppc := PodFunc(h)
	return &ppc
}

// Add some middleware to a particular handler
func (p *PodFunc) Add(m ...interface{}) http.Handler {
	var n http.Handler
	if m != nil {
		stack := toMiddleware(m)
		for i, s := range stack {
			if i == 0 {
				n = s(p)
			} else {
				n = s(n)
			}
		}
	}
	return n
}

func (p PodFunc) Schema(sc ...*Schema) *PodFunc {
	for _, s := range sc {
		for _, m := range *s {
			p = m(p).(PodFunc)
		}
	}
	return &p
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

// Get the interface type and transform to MiddleWare type.
func toMiddleware(m []interface{}) []MiddleWare {
	var stack []MiddleWare
	if len(m) > 0 {
		for _, f := range m {
			switch v := f.(type) {
			case func(http.ResponseWriter, *http.Request):
				stack = append(stack, mutate(http.HandlerFunc(v)))
			case func(http.Handler) http.Handler:
				stack = append(stack, v)
			default:
				fmt.Println("[x] [", reflect.TypeOf(v), "] is not a valid MiddleWare Type.")
			}
		}
	}
	return stack
}

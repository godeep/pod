Pod
=======

## What is pod ?

Pod is a Middleware chaining module, compatible with
every mux who respect the ` http.Handler ` interface.

## Features

- Use Standard ` func (rw http.ResponseWriter, req *http.Request) ` as Middleware.
- Support also ` func (http.Handler) http.Handler ` signature.
- Middleware Chaining.
- Global Middleware declaration.
- Create Middleware Schema
- Standard Lib compatibily
- Compatible with every custom mux,
who respect the ` http.Handler ` interface.

## Example
```go
package main

import "github.com/squiidz/pod"

func main() {
	// Create a new Pod instance, and set some Global Middleware.
	po := pod.NewPod(GlobalMiddleWare)

	// You can also, create a Schema(), which is a stack
	// of MiddleWare for a specific task
	auth := pod.NewSchema(CheckUser, CheckToken, ValidSession)

	// You can pass multiple schema to the same Handler.
	http.Handle("/account", po.Fuse(AccountHandler).Schema(auth).Add(random))

	// Wrap your global middleware with your handler
	http.Handle("/home", po.Fuse(YourHandler))

	// Add some middleware on a specific handler.
	http.Handle("/", po.Fuse(YourOtherHandler).Add(OtherMiddle)) 

	// Start Listening
	http.ListenAndServe(":8080", nil)
}
```

## Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Write Tests!
4. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request

## License
MIT

Pod
=======

## What is pod ?

Pod is a simple Middleware chaining module, compatible with
every mux who are compatible with ` http.Handler `.

## Features

- Middleware Chaining.
- Global Middleware declaration.
- Standard Lib compatibily
- Compatible with every custom mux,
who respect the ` http.Handler ` interface.

## Example
```go
package main

import "github.com/squiidz/pod"

func main() {
	// Create a new Pod instance.
	p := pod.NewPod()

	// Set some Global Middleware
	p.Glob(GlobalMiddleWare)

	// Wrap your global middleware with your handler
	http.Handle("/home", p.Fuse(YourHandler))

	// And add some middleware on precise handler
	http.Handle("/", p.Fuse(YourOtherHandler).Add(OtherMiddle)) 

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

// Command binloader is a HTTP server which builds Go binaries for a variety of platforms.
package main

import (
	"net/http"
)

func main() {
	http.Handle("/", Packager{})

	http.ListenAndServe(":8080", nil)
}

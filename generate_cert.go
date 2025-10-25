package main

import (
	"crud/tlsutil"
	"fmt"
)

// GenerateCertFiles can be run manually (via `go run generate_cert.go`) if you
// want to explicitly generate cert.pem/key.pem. It does not use a main function
// to avoid redeclaration when the package is built normally.
func GenerateCertFiles() {
	cert := "cert.pem"
	key := "key.pem"
	if err := tlsutil.EnsureCert(cert, key); err != nil {
		fmt.Println("failed:", err)
		return
	}
	fmt.Println("generated", cert, key)
}

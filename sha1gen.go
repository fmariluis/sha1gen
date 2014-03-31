/*
This package returns the first n characters of the sha1 digest
of a string, provided by a commandline flag
The resulting string may be used as a reversible password.
*/

package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
)

func main() {

	// command-line flags
	username := flag.String("username", "randomize_me", "Seed string")
	shaLength := flag.Int("len", 6, "Password length")

	flag.Parse()

	h := sha1.New()
	io.WriteString(h, *username)

	sha := h.Sum(nil)
	len := *shaLength / 2
	c := sha[:len]
	fmt.Println("The password is: \n")
	fmt.Printf("%x\n\n", c)
}

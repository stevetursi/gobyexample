package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	p := fmt.Println

	s := "sha1 this string"
	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	p(s)
	fmt.Printf("%x\n", bs)
}

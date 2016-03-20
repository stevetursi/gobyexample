package main

import (
	"fmt"
	r "math/rand"
	"time"
)

func main() {
	p := fmt.Println

	p(r.Intn(100), ",", r.Intn(100))

	p(r.Float64())
	p((r.Float64()*5)+5, ",", (r.Float64()*5)+5)

	s1 := r.NewSource(time.Now().UnixNano())
	r1 := r.New(s1)

	p(r1.Intn(100), ",", r1.Intn(100))

	s2 := r.NewSource(42)
	r2 := r.New(s2)

	p(r2.Intn(100), ",", r2.Intn(100))

	s3 := r.NewSource(42)
	r3 := r.New(s3)

	p(r3.Intn(100), ",", r3.Intn(100))

}

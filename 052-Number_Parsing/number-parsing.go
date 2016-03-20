package main

import (
	"fmt"
	s "strconv"
)

func main() {
	p := fmt.Println

	f, _ := s.ParseFloat("1.234", 64)
	p(f)

	i, _ := s.ParseInt("123", 0, 64)
	p(i)

	d, _ := s.ParseInt("0x1c8", 0, 64)
	p(d)

	u, _ := s.ParseUint("789", 0, 64)
	p(u)

	k, _ := s.Atoi("135")
	p(k)

	_, e := s.Atoi("wat")
	p(e)

}

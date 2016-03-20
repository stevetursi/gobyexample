package main

import (
	"fmt"
	"os"
)

var f = fmt.Printf

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	f("%v\n", p)
	f("%+v\n", p)
	f("%#v\n", p)
	f("%T\n", p)
	f("%t\n", p)

	f("%t\n", true)
	f("%d\n", 123)
	f("%b\n", 14)
	f("%c\n", 33)

	f("%x\n", 456)
	f("%f\n", 78.9)

	f("%e\n", 12340000.0)
	f("%E\n", 12340000.0)

	f("%s\n", "\"String\"")
	f("%q\n", "\"String\"")
	f("%x\n", "hex this")

	f("%p\n", &p)

	f("|%6d|%6d|\n", 12, 345)

	f("|%6.2f|%6.2f|\n", 1.2, 3.45)
	f("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	f("|%6s|%6s|\n", "foo", "b")
	f("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}

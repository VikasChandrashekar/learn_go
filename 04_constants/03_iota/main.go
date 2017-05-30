package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	_ = iota
	d
	e
)

const (
	_ = iota
	f = (iota * 10)
	g = (iota * 20)
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}

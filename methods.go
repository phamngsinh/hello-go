package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area(val int) int {

	return r.width * r.height * val
}
func (r *rect) permin() int {

	return 2*r.width + 2*r.height
}
func validate(val int) int {
	return val
}
func main() {
	m := validate(10)
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area(m))
	fmt.Println("permim:", r.permin())
	rp := &r
	fmt.Println("area:", rp.area(m))
	fmt.Println("permim:", rp.permin())
}

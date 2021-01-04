package main

import "fmt"

func ret(x int) (int, int, int) {
	return x * x, 2 * x, -x
}
func main() {
	fmt.Println(ret(10))
	//tuple assignment
	y1, y2, y3 := ret(20)
	fmt.Println(y1, y2, y3)
	y1, y2 = y2, y1
	fmt.Println(y1, y2, y3)
	x1, x2, x3 := y1*2, y1*y1, -y1
	fmt.Println(x1, x2, x3)

}

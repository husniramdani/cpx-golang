package main

import "fmt"

func main() {
	var x, y int
	ans := 0
	fmt.Scan(&x, &y)
	for x <= y {
		ans++
		x *= 3
		y *= 2
	}
	fmt.Println(ans)
}

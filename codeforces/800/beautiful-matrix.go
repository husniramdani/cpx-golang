package main

import (
	"fmt"
	"math"
)

func main() {
	num := 0
	i := -1
	for num == 0 {
		fmt.Scan(&num)
		i++
	}
	moves := math.Abs(float64(i/5-2)) + math.Abs(float64(i%5-2))
	fmt.Println(moves)
}

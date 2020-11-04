// ░░░░░░░░░░░░░░░SPINDYZEL░░░░░░░░░░░░░░░
// ░░░░░░░░░░▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄░░░░░░░░░
// ░░░░░░░░▄▀░░░░░░░░░░░░▄░░░░░░░▀▄░░░░░░░
// ░░░░░░░░█░░░░░░░░░░░░▄█▄▄░░▄░░░█░▄▄▄░░░
// ░▄▄▄▄▄░░█░░░░░░▀░░░░▀█░░▀▄░░░░░█▀▀░██░░
// ░██▄▀██▄█░░░▄░░░░░░░██░░░░▀▀▀▀▀░░░░██░░
// ░░▀██▄▀██░░░░░░░░▀░██▀░░░░░░░░░░░░░▀██░
// ░░░░▀████░▀░░░░▄░░░██░░░▄█░░░░▄░▄█░░██░
// ░░░░░░░▀█░░░░▄░░░░░██░░░░▄░░░▄░░▄░░░██░
// ░░░░░░░▄█▄░░░░░░░░░░░▀▄░░▀▀▀▀▀▀▀▀░░▄▀░░
// ░░░░░░█▀▀█████████▀▀▀▀████████████▀░░░░
// ░░░░░░████▀░░███▀░░░░░░▀███░░▀██▀░░░░░░

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	ans := n + 1
	find := false
	for !find {
		hitung := make(map[rune]int)
		temp := strconv.Itoa(ans)
		for _, x := range temp {
			hitung[x]++
		}
		if len(hitung) == 4 {
			find = true
			s = temp
		}
		ans++
	}
	fmt.Println(s)
}

// better algorithma
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var y int
// 	fmt.Scan(&y)
// 	for y++; len(map[int]bool{
// 		y % 10:       true,
// 		y / 10 % 10:  true,
// 		y / 100 % 10: true,
// 		y / 1000:     true}) != 4; y++ {
// 	}
// 	fmt.Println(y)
// }

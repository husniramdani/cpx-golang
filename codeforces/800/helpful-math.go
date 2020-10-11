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

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var a, b, c int
	for _, x := range s {
		if x == '1' {
			a++
		} else if x == '2' {
			b++
		} else if x == '3' {
			c++
		}
	}
	for a+b+c > 0 {
		if a > 0 {
			fmt.Print(1)
			a--
		} else if b > 0 {
			fmt.Print(2)
			b--
		} else if c > 0 {
			fmt.Print(3)
			c--
		}
		if a+b+c > 0 {
			fmt.Print("+")
		}
	}
}

// better approach
// package main

// import(
// 	"fmt"
// 	"sort"
// 	"strings"
// )

// func main(){
// 	var s string
// 	fmt.Scan(&s)
// 	ans := strings.Split(s, "+")
// 	sort.Strings(ans)
// 	fmt.Println(strings.Join(ans, "+"))
// }

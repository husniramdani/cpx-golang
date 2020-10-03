package main

import "fmt"

func main() {
	var n, p int
	ans := 0
	limit := 0
	fmt.Scan(&n, &p)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if i+1 == p {
			limit = arr[i]
		}
	}
	for _, x := range arr {
		if x >= limit && x > 0 {
			ans++
		}
	}
	fmt.Println(ans)
}

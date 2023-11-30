//go:build ignore

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
	"bufio"
	"fmt"
	"go-learn/common"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

// adding some changes
// basic logic
func miniMaxSum2(arr []int32) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var min, max int64
	for i, val := range arr {
		if i < 4 {
			min += int64(val)
		}
		if i > 0 {
			max += int64(val)
		}
	}
	fmt.Println(min, max)
}

// medium, sum all and - with min or max
func miniMaxSum(arr []int32) {
	var sum, min, max int64 = 0, math.MaxInt, 0
	for _, val := range arr {
		sum += int64(val)
		if int64(val) < min {
			min = int64(val)
		}
		if int64(val) > max {
			max = int64(val)
		}
	}
	fmt.Println(sum-max, sum-min)
}

// main
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(common.ReadLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		common.CheckError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

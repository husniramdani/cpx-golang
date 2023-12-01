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
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'divisibleSumPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY ar
 */

// func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
// 	// Write your code here
// 	result := 0
// 	for i := 0; i < int(n); i++ {
// 		for j := i + 1; j < int(n); j++ {
// 			sum := ar[i] + ar[j]
// 			if sum%k == 0 {
// 				result++
// 			}
// 		}
// 	}
// 	return int32(result)
// }

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	m := make(map[int32]int32)
	for i := 0; i < len(ar); i++ {
		m[ar[i]%k]++
	}

	var result int32

	if c, ok := m[0]; ok {
		result = c * (c - 1) / 2
	}

	for i := int32(1); i < k/2+1; i++ {
		r := k - i
		cl, okl := m[i]
		cr, okr := m[r]
		if okl && okr {
			if r == i {
				result += cl * (cl - 1) / 2
			} else {
				result += cl * cr
			}
		}
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(common.ReadLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	common.CheckError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	common.CheckError(err)
	k := int32(kTemp)

	arTemp := strings.Split(strings.TrimSpace(common.ReadLine(reader)), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		common.CheckError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := divisibleSumPairs(n, k, ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

//go:build ignore

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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

// updated
func plusMinus(arr []int32) {
	var pos, zero, neg float32
	var arrLen = float32(len(arr))
	for _, val := range arr {
		if val > 0 {
			pos++
		}
		if val == 0 {
			zero++
		}
		if val < 0 {
			neg++
		}
	}
	fmt.Printf("%.6f\n%.6f\n%.6f", pos/arrLen, neg/arrLen, zero/arrLen)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(common.ReadLine(reader)), 10, 64)
	common.CheckError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(common.ReadLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		common.CheckError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

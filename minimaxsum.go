package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func miniMaxSum(arr []int32) {
	var answer []int64
	for i := 0; i < len(arr); i++ {
		var sum int64
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			sum += int64(arr[j])
		}

		if len(answer) == 0 {
			answer = append(answer, sum)
		}

		if sum < answer[0] {
			answer[0] = sum
		}
		if len(answer) == 1 {
			answer = append(answer, sum)
		}
		if sum > answer[1] {
			answer[1] = sum
		}
	}
	fmt.Printf("%d %d", answer[0], answer[1])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

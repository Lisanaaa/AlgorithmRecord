package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumFromOneLine(ir *bufio.Reader) int {
	stream, _ := ir.ReadString('\n')
	// Trims the stream and then splits
	trimmedStream := strings.TrimSpace(stream)
	num, _ := strconv.Atoi(trimmedStream)
	return num
}

func getNumsFromOneLine(ir *bufio.Reader) []int {
	stream, _ := ir.ReadString('\n')
	// Trims the stream and then splits
	trimmedStream := strings.TrimSpace(stream)
	splitArr := strings.Split(trimmedStream, " ")
	// convert strings to integers and store them in a slice
	nums := make([]int, len(splitArr))
	for index, val := range splitArr {
		nums[index], _ = strconv.Atoi(val)
	}
	return nums
}

func main() {
	ir := bufio.NewReader(os.Stdin)
	_ = getNumFromOneLine(ir)
	nums := getNumsFromOneLine(ir)

	pos, neg := 0, 0
	for _, num := range nums {
		if num > 0 {
			pos += 1
		} else if num < 0 {
			neg += 1
		}
	}

	if neg >= (len(nums)+1)/2 {
		fmt.Println(-1)
	} else if pos >= (len(nums)+1)/2 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

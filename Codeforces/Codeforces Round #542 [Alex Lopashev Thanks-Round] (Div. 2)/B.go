package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getNumFromOneLine(ir *bufio.Reader) int{
	stream, _ := ir.ReadString('\n')
	// Trims the stream and then splits
	trimmedStream := strings.TrimSpace(stream)
	num, _ := strconv.Atoi(trimmedStream)
	return num
}

func getNumsFromOneLine(ir *bufio.Reader) []int{
	stream, _ := ir.ReadString('\n')
	// Trims the stream and then splits
	trimmedStream := strings.TrimSpace(stream)
	splitArr := strings.Split(trimmedStream, " ")
	// convert strings to integers and store them in a slice
	nums := make([]int, len(splitArr))
	for index,val := range splitArr{
		nums[index], _ = strconv.Atoi(val)
	}
	return nums
}

func main() {
	ir := bufio.NewReader(os.Stdin)
	n := getNumFromOneLine(ir)
	nums := getNumsFromOneLine(ir)
	//fmt.Println(nums)
	lookup := make(map[int][]int)
	for i, num := range nums {
		lookup[num] = append(lookup[num], i)
	}
	// res最终可能会很大，所以需要用int64
	var res int64
	// 两个人都必须要先到达1, 所以要加上从0到1的距离
	res = int64(lookup[1][0] + lookup[1][1])
	// 每一轮两个人都要从i到i+1，无论这轮他们怎么走，之后的所有轮都是两个人从i+1出发，所以我们只要取得此轮最优即可
	for i := 1; i < n; i += 1 {
		// 一个人从第一个i到第一个i+1, 另一个人从第二个i到第二个i+1
		path1 := math.Abs(float64(lookup[i][0] - lookup[i+1][0])) + math.Abs(float64(lookup[i][1] - lookup[i+1][1]))
		// 一个人从第一个i到第二个i+1, 另一个人从第二个i到第一个i+1
		path2 := math.Abs(float64(lookup[i][0] - lookup[i+1][1])) + math.Abs(float64(lookup[i][1] - lookup[i+1][0]))
		// 加上这一轮的最优距离
		res += int64(math.Min(path1, path2))
	}
	fmt.Println(res)
}





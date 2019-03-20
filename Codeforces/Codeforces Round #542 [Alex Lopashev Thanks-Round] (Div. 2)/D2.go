package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func dist(a, b, n int) int {
	return (b + n - a) % n
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
	nm := getNumsFromOneLine(ir)
	n, m := nm[0], nm[1]
	des := make([]int, 5004) // station i 的 destination 是des[i]
	cnt := make([]int, 5004) // station i 里面共有 candy cnt[1]个
	for i := 0; i < m; i += 1 {
		ab := getNumsFromOneLine(ir)
		a, b := ab[0], ab[1]
		if des[a] == 0 {
			des[a] = b
		} else {
			if dist(a, b, n) < dist(a, des[a], n) {
				des[a] = b // 永远使得最后一轮距离最短
			}
		}
		cnt[a] += 1
	}

	res := make([]string, 0)
	for i := 1; i < n+1; i++ { // 起点是 station i
		curRes := 0
		for j := 1; j < n+1; j++ {
			if des[j] > 0 {
				// curRes 永远取最大那个
				// 对于某个 station j 来说，所需要走的距离为从起点 station i 到 station j 的距离       +
				//                                    (candy 数 - 1) * n                       +
				//                                     从 station j 到达最终的 destination des[j]
				curRes = int(math.Max(float64(curRes), float64(dist(i, j, n) + (cnt[j] - 1) * n + dist(j, des[j], n))))
			}
		}
		res = append(res, strconv.Itoa(curRes))
	}
	fmt.Println(strings.Join(res, " "))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var i int
	var input [1001]string
	var split_input []string
	var N int
	var D int
	var d int
	var d_sum int
	var ans int

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	N, _ = strconv.Atoi(split_input[0])
	D, _ = strconv.Atoi(split_input[1])

	for i = 1; i < N; i++ {
		sc.Scan()
		input[i] = sc.Text()
	}

	for i = 1; i < N; i++ {
		d, _ = strconv.Atoi(input[i])
		d_sum = d_sum + d
	}
	ans = (D * D * N) - (D * d_sum)
	fmt.Println(ans)

}

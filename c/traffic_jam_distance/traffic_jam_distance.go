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
	var input [1001]string
	var input_str []string
	var N int
	var M int
	var tmp int
	var sum int

	i := 0
	for sc.Scan() {
		input[i] = sc.Text()
		i++
	}

	//fmt.Println(sum)

	input_str = strings.Split(input[0], " ")
	N, _ = strconv.Atoi(input_str[0])
	M, _ = strconv.Atoi(input_str[1])

	// fmt.Println(s)
	for i = 1; i < N; i++ {
		tmp, _ = strconv.Atoi(input[i])
		if tmp <= M {
			sum = sum + tmp
		}
	}

	fmt.Println(sum)
}

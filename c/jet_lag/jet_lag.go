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
	var input [101]string
	var split_input []string
	var N int

	var tmp_s int
	var tmp_f int
	var tmp_t int
	var tmp int
	var min int = 48
	var max int

	sc.Scan()
	input[0] = sc.Text()
	N, _ = strconv.Atoi(input[0])

	for i = 0; i < N; i++ {
		sc.Scan()
		input[i] = sc.Text()
		split_input = strings.Split(input[i], " ")
		tmp_s, _ = strconv.Atoi(split_input[0])
		tmp_f, _ = strconv.Atoi(split_input[1])
		tmp_t, _ = strconv.Atoi(split_input[2])
		tmp = tmp_s + tmp_f + (24 - tmp_t)
		//fmt.Println(tmp)
		if tmp > max {
			max = tmp
		}
		if tmp < min {
			min = tmp
		}
	}
	fmt.Println(min)
	fmt.Println(max)

}

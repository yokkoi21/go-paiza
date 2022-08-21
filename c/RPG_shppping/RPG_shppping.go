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
	var input [103]string
	var split_input []string
	var N int
	var T int
	var Q int
	var i int
	var item_value []int
	var cash int

	sc.Scan()
	input[0] = sc.Text()
	N, _ = strconv.Atoi(input[0])

	sc.Scan()
	input[1] = sc.Text()
	split_input = strings.Split(input[1], " ")
	for i = 0; i < N; i++ {
		tmp, _ := strconv.Atoi(split_input[i])
		item_value = append(item_value, tmp)
	}

	sc.Scan()
	input[2] = sc.Text()
	split_input = strings.Split(input[2], " ")
	T, _ = strconv.Atoi(split_input[0])
	Q, _ = strconv.Atoi(split_input[1])

	cash = T
	for i = 0; i < Q; i++ {
		sc.Scan()
		input[i+3] = sc.Text()
		split_input = strings.Split(input[i+3], " ")
		tmp_x, _ := strconv.Atoi(split_input[0])
		tmp_k, _ := strconv.Atoi(split_input[1])
		if item_value[tmp_x-1]*tmp_k < cash {
			cash = cash - (item_value[tmp_x-1] * tmp_k)
		}
	}

	fmt.Println(cash)

}

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
	var input [1002]string
	var split_input []string
	var N int
	var M int
	var K int
	var my_score []string
	var i int
	var j int
	var count int
	var ans_flag []int
	var ans []string
	var ans_str string

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	N, _ = strconv.Atoi(split_input[0])
	M, _ = strconv.Atoi(split_input[1])
	K, _ = strconv.Atoi(split_input[2])

	sc.Scan()
	input[1] = sc.Text()
	my_score = strings.Split(input[1], " ")

	for i = 2; i < M+2; i++ {
		sc.Scan()
		input[i] = sc.Text()
	}

	matrix := make([][]string, M)
	for i = 0; i < M; i++ {
		matrix[i] = make([]string, N)
	}

	for i = 0; i < N; i++ {
		ans_flag = append(ans_flag, 0)
	}

	for i = 0; i < M; i++ {
		count = 0
		split_input = strings.Split(input[i+2], " ")
		for j = 0; j < N; j++ {
			matrix[i][j] = split_input[j]
			if matrix[i][j] == "3" && my_score[j] == "3" {
				count++
			}
		}
		//fmt.Println(count)
		if count >= K {
			for j = 0; j < N; j++ {
				if matrix[i][j] == "3" && my_score[j] == "0" {
					ans_flag[j] = 1
				}
			}
		}
	}

	for i = 0; i < N; i++ {
		if ans_flag[i] == 1 {
			ans = append(ans, strconv.Itoa(i+1))
		}
	}
	if len(ans) == 0 {
		fmt.Println("no")
	} else {
		ans_str = strings.Join(ans, " ")
		fmt.Println(ans_str)
	}

}

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
	var input [11]string
	var split_input []string
	var N int
	var M int
	var i int
	var j int
	var ans []int
	var ans_s []string
	var check bool

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	N, _ = strconv.Atoi(split_input[0])
	M, _ = strconv.Atoi(split_input[1])

	for i = 1; i < N+1; i++ {
		sc.Scan()
		input[i] = sc.Text()
	}

	matrix := make([][]int, N)
	for i := 0; i < N; i++ {
		matrix[i] = make([]int, N)
	}
	for i = 1; i < N+1; i++ {
		split_input = strings.Split(input[i], " ")
		for j = 0; j < len(split_input); j++ {
			tmp, _ := strconv.Atoi(split_input[j])
			matrix[i-1][j] = tmp
			//matrix[i-1] = append(matrix[i-1], tmp)
		}
		//mt.Println(matrix)
	}
	for j = 0; j < N; j++ {
		for i = 0; i < N; i++ {
			if matrix[i][j] >= M {
				//fmt.Println(matrix[i][j])
				check = true
			}
		}
		if check == false {
			//fmt.Println("test")
			ans = append(ans, j+1)
		}
		check = false
	}
	if len(ans) > 0 {
		for i = 0; i < len(ans); i++ {
			ans_s = append(ans_s, strconv.Itoa(ans[i]))
		}
		fmt.Println(strings.Join(ans_s, " "))
	} else {
		fmt.Println("wait")
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input [31]string
	var split_input []string
	var N int
	var i int
	var j int
	var tmp int
	var ans []int

	sc.Scan()
	input[0] = sc.Text()
	N, _ = strconv.Atoi(input[0])

	for i = 1; i < N+1; i++ {
		sc.Scan()
		input[i] = sc.Text()
	}

	matrix := make([][]int, N+2)
	for i = 0; i < N+2; i++ {
		matrix[i] = make([]int, N+2)
	}

	//fmt.Println(matrix)

	for i = 1; i < N+1; i++ {
		split_input = strings.Split(input[i], " ")
		for j = 1; j < N+1; j++ {
			tmp, _ = strconv.Atoi(split_input[j-1])
			matrix[i][j] = tmp
		}
	}

	//fmt.Println(matrix)

	for i = 1; i < N+1; i++ {
		for j = 1; j < N+1; j++ {
			if matrix[i][j] > matrix[i-1][j] && matrix[i][j] > matrix[i][j-1] && matrix[i][j] > matrix[i][j+1] && matrix[i][j] > matrix[i+1][j] {
				ans = append(ans, matrix[i][j])
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ans)))
	for i = 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}

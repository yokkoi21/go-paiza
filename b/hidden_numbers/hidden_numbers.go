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
	var input [10]string
	var split_input []string
	var H int
	var W int
	var K int
	var i int
	var j int
	var l int
	var tmp int
	var max_s string
	var max_i int

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	H, _ = strconv.Atoi(split_input[0])
	W, _ = strconv.Atoi(split_input[1])
	K, _ = strconv.Atoi(split_input[2])

	for i = 1; i < H+1; i++ {
		sc.Scan()
		input[i] = sc.Text()
	}

	matrix := make([][]string, H+2)
	for i = 0; i < H+2; i++ {
		matrix[i] = make([]string, W+2)
	}

	//fmt.Println(matrix)

	for i = 1; i < H+1; i++ {
		split_input = strings.Split(input[i], "")
		for j = 1; j < W+1; j++ {
			matrix[i][j] = split_input[j-1]
		}
	}

	//fmt.Println(matrix)

	for i = 1; i <= H; i++ {
		for j = 1; j <= W-K+1; j++ {
			max_s = ""
			for l = 0; l < K; l++ {
				max_s += matrix[i][j+l]
				//fmt.Println(max_s)
			}
			tmp, _ = strconv.Atoi(max_s)
			if max_i < tmp {
				max_i = tmp
			}
		}
	}
	for i = 1; i <= H-K+2; i++ {
		for j = 1; j <= W; j++ {
			max_s = ""
			for l = 0; l < K; l++ {
				max_s += matrix[i+l][j]
				//fmt.Println(max_s)
			}
			tmp, _ = strconv.Atoi(max_s)
			if max_i < tmp {
				max_i = tmp
			}
		}
	}

	fmt.Println(max_i)

}

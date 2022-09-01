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
	var input [51]string
	var split_input []string
	var H int
	var W int
	var i int
	var j int
	var ans int

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	H, _ = strconv.Atoi(split_input[0])
	W, _ = strconv.Atoi(split_input[1])

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

	for i = 1; i < H+1; i++ {
		for j = 1; j < W+1; j++ {
			tmp := 4
			if matrix[i][j] == "#" {
				if matrix[i-1][j] == "#" {
					tmp--
				}
				if matrix[i][j-1] == "#" {
					tmp--
				}
				if matrix[i+1][j] == "#" {
					tmp--
				}
				if matrix[i][j+1] == "#" {
					tmp--
				}
				ans += tmp
				//fmt.Println(ans)
			}
		}
	}

	fmt.Println(ans)

}

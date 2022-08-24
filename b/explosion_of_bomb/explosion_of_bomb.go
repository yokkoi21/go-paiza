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
	var split_input []string
	var H int
	var W int
	var i int
	var j int
	var k int
	var l int
	var count int

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	H, _ = strconv.Atoi(split_input[0])
	W, _ = strconv.Atoi(split_input[1])

	for i = 0; i < H; i++ {
		sc.Scan()
		input[i] = sc.Text()
	}

	matrix := make([][]string, H)
	for i := 0; i < H; i++ {
		matrix[i] = make([]string, W)
	}
	matrix2 := make([][]string, H)
	for i := 0; i < H; i++ {
		matrix2[i] = make([]string, W)
	}
	for i = 0; i < H; i++ {
		split_input = strings.Split(input[i], "")
		for j = 0; j < W; j++ {
			matrix[i][j] = split_input[j]
			if matrix[i][j] == "#" {
				// fmt.Print(i)
				// fmt.Print(" ")
				// fmt.Println(j)
				for k = 0; k < H; k++ {
					matrix2[k][j] = "#"
				}
				for l = 0; l < W; l++ {
					matrix2[i][l] = "#"
				}
			}
		}
	}
	for i = 0; i < H; i++ {
		for j = 0; j < W; j++ {
			if matrix2[i][j] == "#" {
				count++
			}
		}
	}

	//fmt.Println(matrix2)
	fmt.Println(count)
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input [1002]string
	var split_input []string
	var split_input2 []string
	var H int
	var W int
	var x int
	var y int
	var n int
	var i int

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	W, _ = strconv.Atoi(split_input[0])
	H, _ = strconv.Atoi(split_input[1])
	n, _ = strconv.Atoi(split_input[2])
	sc.Scan()
	input[1] = sc.Text()
	split_input2 = strings.Split(input[1], " ")
	x, _ = strconv.Atoi(split_input2[0])
	y, _ = strconv.Atoi(split_input2[1])

	for i = 0; i < n; i++ {
		sc.Scan()
		input[i] = sc.Text()
	}

	// W -= 1
	// H -= 1

	for i = 0; i < n; i++ {
		split_input = strings.Split(input[i], " ")
		tmp, _ := strconv.Atoi(split_input[1])
		switch split_input[0] {
		case "U":
			y = (y + tmp) % H
		case "D":
			if y-tmp < 0 {
				y = H - (int(math.Abs(float64(y-tmp))) % H)
			} else {
				y -= tmp
			}
		case "R":
			x = (x + tmp) % W
		case "L":
			if x-tmp < 0 {
				x = W - (int(math.Abs(float64(x-tmp))) % W)
			} else {
				x -= tmp
			}
		}
	}

	fmt.Print(x)
	fmt.Print(" ")
	fmt.Println(y)
}

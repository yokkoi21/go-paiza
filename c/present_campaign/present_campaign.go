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
	var input [21]string
	var input_str []string
	var N int
	var X int
	var Y int

	i := 0
	sc.Scan()
	input[0] = sc.Text()

	input_str = strings.Split(input[0], " ")
	N, _ = strconv.Atoi(input_str[0])
	X, _ = strconv.Atoi(input_str[1])
	Y, _ = strconv.Atoi(input_str[2])

	// fmt.Println(s)
	for i = 1; i < N+1; i++ {
		if (i % X) == 0 {
			if (i % Y) == 0 {
				fmt.Println("AB")
			} else {
				fmt.Println("A")
			}
		} else if (i % Y) == 0 {
			if (i % X) == 0 {
				fmt.Println("AB")
			} else {
				fmt.Println("B")
			}
		} else {
			fmt.Println("N")
		}
	}

}

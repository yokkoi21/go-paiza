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
	var input [1001]string
	var N int
	var split_input_pre []string
	var split_input_now []string

	sc.Scan()
	input[0] = sc.Text()
	N, _ = strconv.Atoi(input[0])

	sc.Scan()
	input[1] = sc.Text()
	split_input_pre = strings.Split(input[1], "")

	for i = 2; i < N+1; i++ {
		sc.Scan()
		input[i] = sc.Text()
	}
	//fmt.Println(input)
	for i = 2; i < N; i++ {
		split_input_now = strings.Split(input[i], "")
		// fmt.Println(i)
		// fmt.Println(split_input_pre)
		// fmt.Println(split_input_now)
		if split_input_now[0] != split_input_pre[len(split_input_pre)-1] {
			fmt.Println(split_input_pre[len(split_input_pre)-1] + " " + split_input_now[0])
			break
		} else {
			split_input_pre = split_input_now
			// fmt.Println(split_input_pre)
			// fmt.Println(split_input_now)
		}
		if i == N-1 {
			fmt.Println("Yes")
		}
	}

}

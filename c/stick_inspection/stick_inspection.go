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
	var input [2]string
	var N int
	var M int
	var split_input []string
	var tmp int
	var count int
	var check bool

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	N, _ = strconv.Atoi(split_input[0])
	M, _ = strconv.Atoi(split_input[1])

	sc.Scan()
	input[1] = sc.Text()
	split_input = strings.Split(input[1], " ")

	for i = 0; i < N; i++ {
		tmp, _ = strconv.Atoi(split_input[i])
		//fmt.Println(tmp)
		if tmp == 0 {
			count++
		} else {
			count = 0
		}
		if count >= M {
			check = true
			break
		}
		//fmt.Println(count)
	}
	if check == true {
		fmt.Println("NG")
	} else {
		fmt.Println("OK")
	}

}

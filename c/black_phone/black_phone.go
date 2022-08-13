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
	var input [1]string
	var split_input []string
	var distance [10]int = [10]int{12, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	var tmp int
	var cnt int
,m
	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], "")
	//fmt.Println(split_input)

	for i = 0; i < len(split_input); i++ {
		if split_input[i] == "-" {
			continue
		} else {
			tmp, _ = strconv.Atoi(split_input[i])
			//fmt.Println(distance[tmp])
			cnt = cnt + distance[tmp]*2
		}
	}
	fmt.Println(cnt)

}

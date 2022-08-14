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
	var input_split []string
	var X_s []string
	var Y_s []string
	var tmp int
	var tmp_X_i int
	var tmp_Y_i int
	var tmp_X_s []string
	var tmp_Y_s []string

	sc.Scan()
	input[0] = sc.Text()
	input_split = strings.Split(input[0], " ")
	//fmt.Println(input_split)
	X_s = strings.Split(input_split[0], "")
	Y_s = strings.Split(input_split[1], "")
	//fmt.Println(X_s)
	//fmt.Println(Y_s)
	for i = 0; i < len(X_s); i++ {
		tmp, _ = strconv.Atoi(X_s[i])
		tmp_X_i = tmp_X_i + tmp
	}
	tmp_X_s = strings.Split(strconv.Itoa(tmp_X_i), "")
	for i = 0; i < len(Y_s); i++ {
		tmp, _ = strconv.Atoi(Y_s[i])
		tmp_Y_i = tmp_Y_i + tmp
	}
	tmp_Y_s = strings.Split(strconv.Itoa(tmp_Y_i), "")

	if len(tmp_X_s) > 1 && len(tmp_Y_s) > 1 {
		if tmp_X_s[1] < tmp_Y_s[1] {
			fmt.Println("Alice")
		} else if tmp_X_s[1] > tmp_Y_s[1] {
			fmt.Println("Bob")
		} else {
			fmt.Println("Draw")
		}
	} else if len(tmp_X_s) == 1 && len(tmp_Y_s) > 1 {
		if tmp_X_s[0] < tmp_Y_s[1] {
			fmt.Println("Alice")
		} else if tmp_X_s[0] > tmp_Y_s[1] {
			fmt.Println("Bob")
		} else {
			fmt.Println("Draw")
		}
	} else if len(tmp_X_s) > 1 && len(tmp_Y_s) == 1 {
		if tmp_X_s[1] < tmp_Y_s[0] {
			fmt.Println("Alice")
		} else if tmp_X_s[1] > tmp_Y_s[0] {
			fmt.Println("Bob")
		} else {
			fmt.Println("Draw")
		}
	} else if len(tmp_X_s) == 1 && len(tmp_Y_s) == 1 {
		if tmp_X_s[0] < tmp_Y_s[0] {
			fmt.Println("Alice")
		} else if tmp_X_s[0] > tmp_Y_s[0] {
			fmt.Println("Bob")
		} else {
			fmt.Println("Draw")
		}
	}

}

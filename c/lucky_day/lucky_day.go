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
	var i_str string
	var input [1]string
	var cnt int

	sc.Scan()
	input[0] = sc.Text()

	for i = 0; i < 366; i++ {
		i_str = strconv.Itoa(i)
		if strings.Contains(i_str, input[0]) {
			//fmt.Println(i_str)
			cnt++
		}
	}
	fmt.Println(cnt)

}

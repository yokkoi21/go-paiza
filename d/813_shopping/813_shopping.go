package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input [1]string
	var cash int

	sc.Scan()
	input[0] = sc.Text()
	cash, _ = strconv.Atoi(input[0])

	if cash >= 1000 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

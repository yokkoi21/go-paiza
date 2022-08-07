package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input [101]string
	var s []string

	i := 0
	sc.Scan()
	input[0] = sc.Text()

	s = strings.Split(input[0], "")
	s = append([]string{"+"}, s[0:]...)
	s = append(s, "+")
	// fmt.Println(s)
	for i = 0; i < len(s); i++ {
		fmt.Print("+")
	}
	fmt.Print("\n")
	for i = 0; i < len(s); i++ {
		fmt.Print(s[i])
	}
	fmt.Print("\n")
	for i = 0; i < len(s); i++ {
		fmt.Print("+")
	}
	fmt.Print("\n")

}

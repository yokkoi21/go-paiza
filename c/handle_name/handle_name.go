package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var input [21]string
	var s []string
	var vowels [10]string = [10]string{"a", "i", "u", "e", "o", "A", "I", "U", "E", "O"}
	var ans []string

	i := 0
	j := 0
	sc.Scan()
	input[0] = sc.Text()

	s = strings.Split(input[0], "")
	// fmt.Println(s)
	for i = 0; i < len(s); i++ {
		for j = 0; j < 10; j++ {
			if strings.Contains(s[i], vowels[j]) == true {
				break
			}
			if j == 9 {
				//fmt.Println(" s:" + s[i] + " vowels:" + vowels[j])
				ans = append(ans, s[i])
			}
		}

	}
	fmt.Println(strings.Join(ans, ""))

}

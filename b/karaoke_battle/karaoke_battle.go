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
	var input [11]string
	var split_input []string
	var N int
	var M int
	var i int
	var j int
	var tmp_score int
	var max_score int
	var true_step []float64

	sc.Scan()
	input[0] = sc.Text()
	split_input = strings.Split(input[0], " ")
	N, _ = strconv.Atoi(split_input[0])
	M, _ = strconv.Atoi(split_input[1])

	for i = 0; i < M; i++ {
		sc.Scan()
		tmp_s := sc.Text()
		tmp_f, _ := strconv.ParseFloat(tmp_s, 64)
		true_step = append(true_step, tmp_f)
	}

	for i = 0; i < N; i++ {
		tmp_score = 100
		for j = 0; j < M; j++ {
			sc.Scan()
			tmp_s := sc.Text()
			tmp_f, _ := strconv.ParseFloat(tmp_s, 64)
			if math.Abs(true_step[j]-tmp_f) <= 5 {
				continue
			} else if math.Abs(true_step[j]-tmp_f) <= 10 {
				tmp_score = tmp_score - 1
			} else if math.Abs(true_step[j]-tmp_f) <= 20 {
				tmp_score = tmp_score - 2
			} else if math.Abs(true_step[j]-tmp_f) <= 30 {
				tmp_score = tmp_score - 3
			} else {
				tmp_score = tmp_score - 5
			}
		}
		if tmp_score >= max_score {
			max_score = tmp_score
		}
	}

	fmt.Println(max_score)

}

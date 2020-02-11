package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [20]string
	var split_input []string
	var n int
	var i,j int
	var operator [11]string
	var num [11]int
	var max,min int
	var ans_check int
	

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ = strconv.Atoi(input[0])


	min = 0
	max = 999999999999999999
	for i = 0;i < n;i++{
		split_input = strings.Split(input[i+1]," ")
		operator[i] = split_input[0]
		num[i],_ = strconv.Atoi(split_input[1])
		if operator[i] == "<"{
			if max > num[i]{
				max = num[i]
			}
		}else if operator[i] == ">"{
			if min < num[i]{
				min = num[i]
			}
		}
	}

	if max != 0 && min != 0{
		for i = min;i < max;i++{
			for j = 0;j < n;j++{
				if operator[j] == "<"{
					if i < num[j]{
						ans_check += 1
					}
				}else if operator[j] == ">"{
					if i > num[j]{
						ans_check += 1
					}
				}else {
					if i % num[j] == 0{
						ans_check += 1
					}
				}
			}
			if ans_check == n{
				fmt.Println(i)
			}
			ans_check = 0
		}
	}else if max != 0{
		for i = 0;i < max;i++{
			for j = 0;j < n;j++{
				if operator[j] == "<"{
					if i < num[j]{
						ans_check += 1
					}
				}else if operator[j] == ">"{
					if i > num[j]{
						ans_check += 1
					}
				}else {
					if i % num[j] == 0{
						ans_check += 1
					}
				}
			}
			if ans_check == n{
				fmt.Println(i)
			}
			ans_check = 0
		}
	}else if min != 0{
		for i = min;i < 999999999999999999;i++{
			for j = 0;j < n;j++{
				if operator[j] == "<"{
					if i < num[j]{
						ans_check += 1
					}
				}else if operator[j] == ">"{
					if i > num[j]{
						ans_check += 1
					}
				}else {
					if i % num[j] == 0{
						ans_check += 1
					}
				}
			}
			if ans_check == n{
				fmt.Println(i)
			}
			ans_check = 0
		}
	}else{
		for i = 0;i < 999999999999999999;i++{
			for j = 0;j < n;j++{
				if operator[j] == "<"{
					if i < num[j]{
						ans_check += 1
					}
				}else if operator[j] == ">"{
					if i > num[j]{
						ans_check += 1
					}
				}else {
					if i % num[j] == 0{
						ans_check += 1
					}
				}
			}
			if ans_check == n{
				fmt.Println(i)
			}
			ans_check = 0
		}
	}
	

	
}
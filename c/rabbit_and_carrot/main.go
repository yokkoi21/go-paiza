package main
import "strings"
import "os"
import "strconv"
import "fmt"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [102]string
	var split_input []string
	var n,s,p,i int
	var tmp_m,tmp_n int
	var ans [100]int
	var ans_int int
	var max int

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	n,_ = strconv.Atoi(split_input[0])
	s,_ = strconv.Atoi(split_input[1])
	p,_ = strconv.Atoi(split_input[2])
	
	j := 0
	for i = 1;i < n+1;i++{
		split_input = strings.Split(input[i]," ")
		tmp_m,_ = strconv.Atoi(split_input[0])
		tmp_n,_ = strconv.Atoi(split_input[1])
		if tmp_n <= s + p && tmp_n >= s - p{
			if max < tmp_m {
				ans_int = i
				max = tmp_m
			}
			ans[j] = i
			j++
		}
	}

	count := 0
	for i = 0;i < len(ans);i++{
		if ans[i] != 0{
			count += 1
		}
	}
	if count == 0{
		fmt.Println("not found")	
	}else{
		fmt.Println(ans_int)
	}
	
}
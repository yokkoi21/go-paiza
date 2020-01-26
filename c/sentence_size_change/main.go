package main
import "strings"
import "os"
import "strconv"
import "fmt"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [101]string
	var split_input []string
	var h,w,x,i,j int
	var ans string
	var ans_split []string
	var tmp int
	var print_ans [500]string
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	h,_ = strconv.Atoi(split_input[0])
	w,_ = strconv.Atoi(split_input[1])
	x,_ = strconv.Atoi(split_input[2])
	
	for i = 0;i < h;i++{
		ans += input[i+1]
	}
	if ((h * w) % x) == 0 {
		tmp = (h * w) / x
	}else {
		tmp = ((h * w) / x) + 1
	}
	// fmt.Println(ans)
	// fmt.Println(tmp)
	ans_split = strings.Split(ans,"")
	for i = 0;i < tmp;i++{
		if ((h * w) % x) != 0 && i == tmp - 1{
			for j = 0;j < ((h * w) % x);j++{
				print_ans[i] += ans_split[(i * x) + j]
			}
		}else {
			for j = 0;j < x;j++{
				print_ans[i] += ans_split[(i * x) + j]
			}
		}
		
		fmt.Println(print_ans[i])
	}
	
	
}
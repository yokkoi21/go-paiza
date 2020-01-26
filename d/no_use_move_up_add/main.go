package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input string
	var split_input []string
	var num1,num2 []string
	var num1_int,num2_int [3]int
	var ans_tmp [3]int
	var pre_ans [3]string
	var ans string
	var i int

	sc.Scan()
	input = sc.Text()
	split_input = strings.Split(input," ")
	num1 = strings.Split(split_input[0],"")
	num2 = strings.Split(split_input[1],"")
	for i = 0;i < len(num1); i++{
		num1_int[i],_ = strconv.Atoi(num1[i])
	}
	for i = 0;i < len(num2); i++{
		num2_int[i],_ = strconv.Atoi(num2[i])
	}
	for i = 0;i < 3; i++{
		if len(num2) == 3 && len(num1) == 1  {
			ans_tmp[0] = num2_int[0]
			ans_tmp[1] = num2_int[1]
			ans_tmp[2] = (num1_int[0] + num2_int[2]) % 10
		}else if len(num1) == 3 && len(num2) == 1 {
			ans_tmp[0] = num1_int[0]
			ans_tmp[1] = num1_int[1]
			ans_tmp[2] = (num1_int[2] + num2_int[0]) % 10
		}else if len(num2) == 2 && len(num1) == 1  {
			ans_tmp[0] = num2_int[0]
			ans_tmp[1] = (num1_int[0] + num2_int[1]) % 10
		}else if len(num1) == 2 && len(num2) == 1  {
			ans_tmp[0] = num1_int[0]
			ans_tmp[1] = (num1_int[1] + num2_int[0]) % 10
		}else if len(num1) == 3 && len(num2) == 2 {
			ans_tmp[0] = num1_int[0]
			ans_tmp[1] = (num1_int[1] + num2_int[0]) % 10
			ans_tmp[2] = (num1_int[2] + num2_int[1]) % 10
		}else if len(num2) == 3 && len(num1) == 2 {
			ans_tmp[0] = num2_int[0]
			ans_tmp[1] = (num1_int[0] + num2_int[1]) % 10
			ans_tmp[2] = (num1_int[1] + num2_int[2]) % 10
		}else {
			ans_tmp[i] = (num1_int[i] + num2_int[i]) % 10
		}
		pre_ans[i] = strconv.Itoa(ans_tmp[i])
	}
	if (len(num1) == 2 || len(num1) == 1) && (len(num2) == 2 || len(num2) == 1){
		ans = pre_ans[0] + pre_ans[1]
	}else if len(num1) ==  1 && len(num2) == 1{
		ans = pre_ans[0] 
	}else {
		ans = pre_ans[0] + pre_ans[1] + pre_ans[2]
	}
	
	fmt.Println(ans)

}
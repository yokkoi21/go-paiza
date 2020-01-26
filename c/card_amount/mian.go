package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [101]string
	var split_input []string
	var n,x10_check,max_value,tmp_value,clear_check,ans,i int
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	
	n,_ = strconv.Atoi(input[0])
	split_input = strings.Split(input[1]," ")
	
	for i = 0;i < n;i++{
		if split_input[i] == "x10"{
			x10_check = 1
		}else {
			tmp_value,_ = strconv.Atoi(split_input[i])
			if tmp_value == 0 {
				clear_check = 1
			}else {
				ans += tmp_value
				if max_value < tmp_value{
					max_value = tmp_value
				}
			}
		}
	} 
	if clear_check == 1{
		ans = ans - max_value
	}
	if x10_check == 1{
		ans = ans * 10
	}
	
	
	fmt.Println(ans)
	
}
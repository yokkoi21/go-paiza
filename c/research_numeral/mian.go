package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [33]string
	var split_input []string
	var i,j int
	var n,x int
	var x_b_str string
	var x_b_slice []string
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	n,_ = strconv.Atoi(split_input[0])
	x,_ = strconv.Atoi(split_input[1])
	x_b_str = fmt.Sprintf("%b",x)
	x_b_slice = strings.Split(x_b_str,"")
	//fmt.Println(x_b_slice)
	for i = 1; i < n+1;i++{
		for j = len(x_b_slice)+1;j > 0;j--{
			tmp,_ := strconv.Atoi(input[i])
			if j == tmp{
				tmp2 := len(x_b_slice) - tmp
				fmt.Println(x_b_slice[tmp2])
			}
		}
	}
	
}
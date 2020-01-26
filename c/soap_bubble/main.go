package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [102]string
	var split_input []string
	var i,t,x_tmp,y_tmp,x_max int
	var x,y int
	
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	t,_ = strconv.Atoi(split_input[0])
	x,_ = strconv.Atoi(split_input[1])
	y,_ = strconv.Atoi(split_input[2])
	x_max = x
	x_tmp = x
	y_tmp = y

	for i = 1;i < t+1; i++{
		split_input = strings.Split(input[i]," ")
		x,_ = strconv.Atoi(split_input[0])
		y,_ = strconv.Atoi(split_input[1])
		// fmt.Println(x)
		// fmt.Println(y)
		x_tmp += x
		y_tmp += y
		if x_tmp > x_max {
			x_max = x_tmp
		}
		if y_tmp <= 0{
			break
		}
		
	}
	fmt.Println(x_max)
	
}
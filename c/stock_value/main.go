package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [1001]string
	var split_input []string
	var i int
	var s,e,h,l int
	var start_value,end_value,highest_value,lowest_value int
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ := strconv.Atoi(input[0])
	n += 1
	for i = 1;i < n;i++{
		split_input = strings.Split(input[i]," ")
		s,_ = strconv.Atoi(split_input[0])
		e,_ = strconv.Atoi(split_input[1])
		h,_ = strconv.Atoi(split_input[2])
		l,_ = strconv.Atoi(split_input[3])
		if i == 1{
			start_value = s
			highest_value = h
			lowest_value = l
		}
		if i == (n -1){
			end_value = e
		}
		if h > highest_value{
			highest_value = h
		}
		if l < lowest_value{
			lowest_value = l
		}
	}
	fmt.Printf("%d %d %d %d",start_value,end_value,highest_value,lowest_value)
	
}
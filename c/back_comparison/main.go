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
	var split_split_input []string
	var i,n int
	var tmp1,tmp2 int
	var min,max,tmp float64
	var char string
	
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ = strconv.Atoi(input[0])
	
	for i = 1;i < n+1; i++{
		split_input = strings.Split(input[i]," ")
		char = split_input[0]
		split_split_input = strings.Split(split_input[1],".")
		tmp1,_ = strconv.Atoi(split_split_input[0])
		tmp2,_ = strconv.Atoi(split_split_input[1])
		tmp = float64(tmp1) + float64(tmp2 / 10)
		
		if char == "le" && max > tmp{
			max = tmp
		}else if char == "ge" && min < tmp{
			min = tmp
		}
	}
	
	fmt.Printf("%f %f\n",min,max)
}
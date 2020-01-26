package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [1002]string
	var split_input []string
	var i,a,b,r,n int
	var x,y int
	
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	a,_ = strconv.Atoi(split_input[0])
	b,_ = strconv.Atoi(split_input[1])
	r,_ = strconv.Atoi(split_input[2])
	n,_ = strconv.Atoi(input[1])
	
	for i = 2;i < n+2; i++{
		split_input = strings.Split(input[i]," ")
		x,_ = strconv.Atoi(split_input[0])
		y,_ = strconv.Atoi(split_input[1])
		if ((x - a) * (x - a)) + ((y - b) * (y - b)) >= r*r {
			fmt.Println("silent")
		}else {
			fmt.Println("noisy")
		}
	}
	
}
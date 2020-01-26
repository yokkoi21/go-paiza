package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [42]string
	var split_input []string
	var i,n int
	var p1,p2 int
	var c1,c2 int
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	p1,_ = strconv.Atoi(split_input[0])
	p2,_ = strconv.Atoi(split_input[1])
	n,_ = strconv.Atoi(input[1])
	for i = 2;i < n + 2;i++ {
		split_input = strings.Split(input[i]," ")
		c1,_ = strconv.Atoi(split_input[0])
		c2,_ = strconv.Atoi(split_input[1])
		if p1 < c1 {
			fmt.Println("Low")
		}else if p1 > c1{
			fmt.Println("High")
		}else{
			if p2 > c2 {
				fmt.Println("Low")
			}else {
				fmt.Println("High")
			}
		}
	}
	
	
}
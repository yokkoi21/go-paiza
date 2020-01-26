package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [31]string
	var split_input []string
	var n,a,b,x,i int
	var count [61]int
	
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ = strconv.Atoi(input[0])
	
	for i = 1;i < n+1; i++{
		split_input = strings.Split(input[i]," ")
		a,_ = strconv.Atoi(split_input[0])
		b,_ = strconv.Atoi(split_input[1])
		x = a + b
		count[x] += 1
	}
	max := 0
	for i = 1;i < len(count); i++{
		if count[i] > max {
			max = count[i]
		}
	}
	for i = 1;i < len(count); i++{
		if count[i] == max {
			fmt.Println(i)
			break		
		}
	}
	
	
}
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
	var n,h,w,d,i,r,diameter int
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	
	n,_ = strconv.Atoi(input[0])
	split_input = strings.Split(input[0]," ")
	n,_ = strconv.Atoi(split_input[0])
	r,_ = strconv.Atoi(split_input[1])
	diameter = r * 2
	
	for i = 1;i < n+1;i++{
		split_input = strings.Split(input[i]," ")
		h,_ = strconv.Atoi(split_input[0])
		w,_ = strconv.Atoi(split_input[1])
		d,_ = strconv.Atoi(split_input[2])
		if h >= diameter && w >= diameter && d >= diameter{
			fmt.Println(i)
		}
	} 
	
}
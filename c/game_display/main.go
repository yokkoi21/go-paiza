package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [60]string
	var split_input []string
	var h,w,i int
	var dy,dx int
	var ans int
	
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	h,_ = strconv.Atoi(split_input[0])
	w,_ = strconv.Atoi(split_input[1])
	split_input = strings.Split(input[1]," ")
	dy,_ = strconv.Atoi(split_input[0])
	dx,_ = strconv.Atoi(split_input[1])
	
	if dy < 0{
		dy = dy * (-1)
	}
	if dx < 0{
		dx = dx * (-1)
	}
	ans = ((dy * w) + (dx * h)) - (dy * dx)

	fmt.Println(ans)
	
}
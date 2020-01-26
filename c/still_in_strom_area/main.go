package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [191]string
	var n,i int
	var split_input []string
	var xc,yc,r_1,r_2 int
	var x,y int
	
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	xc,_ = strconv.Atoi(split_input[0])
	yc,_ = strconv.Atoi(split_input[1])
	r_1,_ = strconv.Atoi(split_input[2])
	r_2,_ = strconv.Atoi(split_input[3])
	n,_ = strconv.Atoi(input[1])

	

	for i = 2;i < n+2;i++{
		split_input = strings.Split(input[i]," ")
		x,_ = strconv.Atoi(split_input[0])
		y,_ = strconv.Atoi(split_input[1])
		if (r_1 * r_1) <= ((x - xc) * (x - xc)) + ((y - yc) * (y - yc)) {
			if  (r_2 * r_2) >= ((x - xc) * (x - xc)) + ((y - yc) * (y - yc)) {
				fmt.Println("yes")	
			}else {
				fmt.Println("no")
			}
		} else {
			fmt.Println("no")
		}
	}

	
	
}	
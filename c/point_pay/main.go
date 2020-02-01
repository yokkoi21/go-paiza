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
	var day,price int
	var point_sum int
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ := strconv.Atoi(input[0])
	for i = 1;i < n+1;i++{
		split_input = strings.Split(input[i]," ")
		day,_ = strconv.Atoi(split_input[0])
		price,_ = strconv.Atoi(split_input[1])
		if day == 3 || day == 13 || day ==23 || day == 30 || day == 31{
			point_sum += (price * 3) / 100
		}else if day == 5 || day == 15 || day ==25 {
			point_sum += (price * 5) / 100
		}else {
			point_sum += (price * 1) / 100
		}
	}
	fmt.Println(point_sum)
	
}
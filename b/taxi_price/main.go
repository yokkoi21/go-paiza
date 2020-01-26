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
	var n,x int
	var i int
	var f_dist,f_price,a_dist,a_price int
	var max_price,min_price,tmp_price,tmp_dist int
	
	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	
	min_price = 5000000000
	split_input = strings.Split(input[0]," ")
	n,_ = strconv.Atoi(split_input[0])
	x,_ = strconv.Atoi(split_input[1])

	for i = 1;i < n+1;i++{
		tmp_dist = x
		tmp_price = 0
		split_input = strings.Split(input[i]," ")
		f_dist,_ = strconv.Atoi(split_input[0])
		f_price,_ = strconv.Atoi(split_input[1])
		a_dist,_ = strconv.Atoi(split_input[2])
		a_price,_ = strconv.Atoi(split_input[3])
		tmp_dist -= f_dist
		tmp_price += f_price
		if tmp_dist >= 0{
			tmp_price += ((tmp_dist / a_dist) + 1) * a_price
		}
		
		
		if max_price < tmp_price {
			max_price = tmp_price
		}
		if min_price > tmp_price {
			min_price = tmp_price
		}

	}

	fmt.Printf("%d %d\n",min_price,max_price)

}
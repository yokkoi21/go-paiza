package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input string
	var split_input []string
	var i int
	var coffee_price,d_rate int
	var coffee_price_count int
	

	sc.Scan()
	input = sc.Text()
	split_input = strings.Split(input," ")
	coffee_price,_ = strconv.Atoi(split_input[0])
	d_rate,_ = strconv.Atoi(split_input[1])
	
	for i = 0;i < 10000; i++{
		if coffee_price > 0{
			coffee_price_count += coffee_price
			coffee_price = (coffee_price * (100 - d_rate)) / 100
		}else{
			
		}
	}

	fmt.Println(coffee_price_count)
	
}
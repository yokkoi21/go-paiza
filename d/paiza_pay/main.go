package main
import "fmt"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var tmp string
	var current_money int
	sc.Scan()
	tmp = sc.Text()
	current_money,_ = strconv.Atoi(tmp)
	if current_money <= 10000 {
		current_money += 10000
	}

	fmt.Println(current_money)
}
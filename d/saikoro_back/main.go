package main
import "fmt"
import "strconv"
import "os"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var num int
	var ans int
	sc.Scan()
	num,_ = strconv.Atoi(sc.Text())
	
	ans = 7 - num

	fmt.Println(ans)
}
package main
import "fmt"
import "strconv"
import "os"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var c1,c2 int
	var ans int
	sc.Scan()
	c1,_ = strconv.Atoi(sc.Text())
	sc.Scan()
	c2,_ = strconv.Atoi(sc.Text())
	ans = 180 - (c1 + c2)

	fmt.Println(ans)
}
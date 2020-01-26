package main
import "fmt"
import "os"
import "strconv"
import "bufio"

func main() {
	var n,i int
	var sc = bufio.NewScanner(os.Stdin)
	var ans int
	sc.Scan()
	n,_ = strconv.Atoi(sc.Text())
	for i = 0;i <= n;i++{
		ans += i
	}
	fmt.Println(ans)

}

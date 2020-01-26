package main
import "fmt"
import "strings"
import "strconv"
import "os"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var arr []string
	var ans int
	var line string
	sc.Scan()
	line = sc.Text()
	arr = strings.Split(line," ")
	c,_ := strconv.Atoi(arr[0])
	t,_ := strconv.Atoi(arr[1])
	a,_ := strconv.Atoi(arr[2])
	ans = c + t + a
	fmt.Println(ans)
}
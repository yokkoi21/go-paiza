package main
import "fmt"
import "os"
import "bufio"
import "strconv"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var arr [3]string
	var a,b,x int
	var ans int
	var tani,before,after int
	var i int
	i = 0
	for sc.Scan(){
		arr[i] = sc.Text()
		i++
	}
	x,_ = strconv.Atoi(arr[0])
	a,_ = strconv.Atoi(arr[1])
	b,_ = strconv.Atoi(arr[2])
	tani = x / a 
	before = a * tani
	after = b * tani

	ans = after - before

	fmt.Println(ans)

}
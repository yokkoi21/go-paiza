package main
import "fmt"
import "strconv"
import "os"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var slice [100]int
	var ans int
	i := 0
	for sc.Scan(){
		slice[i],_ = strconv.Atoi(sc.Text())
		i += 1
	}
	floor := 1

	for i = 1; i < slice[0]+1; i++{
		if floor > slice[i]{
			ans += floor - slice[i]
			floor = slice[i]
		} else {
			ans += slice[i] - floor
			floor = slice[i]
		}
	}

	fmt.Println(ans)
}
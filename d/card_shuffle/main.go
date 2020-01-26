package main
import "fmt"
import "strings"
import "strconv"
import "os"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var slice []int
	var ans int
	var line string
	sc.Scan()
	line,_ = strconv.Atoi(sc.Text())
	arr = strings.split(line," ")
	//みかん
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
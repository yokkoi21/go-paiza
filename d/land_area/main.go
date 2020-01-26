package main
import "fmt"
import "strings"
import "strconv"
import "os"
import "bufio"
import "math"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    var num_slice []string
    var sc = bufio.NewScanner(os.Stdin)
	var text string
	var tmp1 int
	var tmp2 float64
    var a,b,c,d int
    sc.Scan()
    text  = sc.Text()
    num_slice = strings.Split(text, " ")
    a, _ = strconv.Atoi(num_slice[0])
    b, _ = strconv.Atoi(num_slice[1])
    c, _ = strconv.Atoi(num_slice[2])
	d, _ = strconv.Atoi(num_slice[3])
	tmp1 = (a*d) - (b*c)
	tmp2 = math.Abs(float64(tmp1)) /2
	ans := int(math.Ceil(tmp2))
    fmt.Println(ans)
    
}
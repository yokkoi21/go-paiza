package main
import "fmt"
import "strconv"
import "os"
import "bufio"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    var sc = bufio.NewScanner(os.Stdin)
    var num1,num2,num3,num4,sum,ans int
    sc.Scan()
    num1, _  = strconv.Atoi(sc.Text())
	sc.Scan()
	num2, _  = strconv.Atoi(sc.Text())
	sc.Scan()
    num3, _  = strconv.Atoi(sc.Text())
	sc.Scan()
    num4, _  = strconv.Atoi(sc.Text())
	sum = num1 + num2 + num3 + num4

	ans = 15 - sum
	fmt.Println(ans)
    
}
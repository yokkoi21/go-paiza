package main
import "fmt"
import "strconv"
import "os"
import "bufio"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    var sc = bufio.NewScanner(os.Stdin)
    var n int 
    sc.Scan()
	n,_ = strconv.Atoi(sc.Text())
	ans := 100 + (n * 10)
	fmt.Println(ans)
    
}
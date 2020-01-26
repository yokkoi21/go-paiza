package main
import "fmt"
import "os"
import "bufio"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    var sc = bufio.NewScanner(os.Stdin)
    var local,domain string
    sc.Scan()
    local  = sc.Text()
	sc.Scan()
	domain  = sc.Text()
	ans := local
	ans += "@"
	ans += domain
	fmt.Println(ans)
    
}
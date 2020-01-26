package main
import "fmt"
import "strconv"
import "os"
import "bufio"

func s_to_m() int{
    // 自分の得意な言語で
    // Let's チャレンジ！！
    var sc = bufio.NewScanner(os.Stdin)
    var n,ans int
    sc.Scan()
    n, _  = strconv.Atoi(sc.Text())
    ans = n * 60
    return ans
}

func main(){
    var answer int
    answer = s_to_m()
    fmt.Println(answer)
}
package main
import "fmt"
import "strings"
import "strconv"
import "os"
import "bufio"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    var num_slice []string
    var sc = bufio.NewScanner(os.Stdin)
    var text string
    var m,n int
    sc.Scan()
    text  = sc.Text()
    num_slice = strings.Split(text, " ")
    m, _ = strconv.Atoi(num_slice[0])
    n, _ = strconv.Atoi(num_slice[1])
    ans := m - n
    fmt.Println(ans)
    
}
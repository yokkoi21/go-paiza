package main
import "fmt"
import "strings"
import "strconv"
import "os"
import "bufio"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    var sc = bufio.NewScanner(os.Stdin)
	var arr [11]string
	var text_slice[] string
	var variable1,variable2 int
	i1 := 0
    for sc.Scan(){
		arr[i1] = sc.Text()
		i1 += 1
	}
	var n,i2 int
	n,_ = strconv.Atoi(arr[0])
	for i2 = 1;i2 <= n; i2++{
		text_slice = strings.Split(arr[i2], " ")
		if (text_slice[0] == "SET") == true {
			i,_ := strconv.Atoi(text_slice[1])
			if  i == 1 {
				variable1,_ = strconv.Atoi(text_slice[2])
			} else if i == 2 {
				variable2,_ = strconv.Atoi(text_slice[2])
			}
		} else if (text_slice[0] == "ADD") == true {
			tmp,_ := strconv.Atoi(text_slice[1])
			variable2 = variable1 + tmp
		} else if (text_slice[0] == "SUB") == true {
			tmp,_ := strconv.Atoi(text_slice[1])
			variable2 = variable1 - tmp
		}
	}
	
	fmt.Printf("%d %d",variable1,variable2)
    
}
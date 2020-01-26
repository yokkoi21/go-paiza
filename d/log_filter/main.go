package main
import "fmt"
import "strings"
import "os"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var arr1 [102]string
	var g string
	i := 0
	for sc.Scan(){
		arr1[i] = sc.Text()
		i++
	}
	g = arr1[1]
	var num int
	num = 0
	for i = 0; i <= 101; i++{
		if i == 0 || i == 1 {
		}else {
			if strings.Contains(arr1[i], g) {
				fmt.Println(arr1[i])
				num = num + 1
			}
		} 
	}
	if num == 0{
		fmt.Println("None")
	}

}
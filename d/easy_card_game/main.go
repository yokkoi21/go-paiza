package main
import "fmt"
import "strings"
import "os"
import "bufio"
import "strconv"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var arr [401]string
	var tmp string
	var line []string
	var num [5]int
	i := 0
	l := 0
	for sc.Scan(){
		arr[i] = sc.Text()
		i++
	}
	
	for i = 1; i <= 401; i++{
		tmp = arr[i]		
		line = strings.Split(tmp,"")
		for l = 0; l < 4; l++{
			num[l],_ = strconv.Atoi(line[l])
			
		}
		if num[0] == num[3] {
			fmt.Println("Four Card")
		} else if num[0] == num[2] || num[1] == num[3]{
			fmt.Println("Three Card")
		} else if num[0] == num[1] && num[2] == num[3]{
			fmt.Println("Two Pair")
		} else if num[0] == num[1] || num[1] == num[2] || num[2] == num[3] {
			fmt.Println("One Pair")
		} else {
			fmt.Println("No Pair")
		}
	}
	

}
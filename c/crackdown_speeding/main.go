package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [1001]string
	var split_input []string
	var n,v int
	var t,p [1001]int
	var i,j int
	var ans string
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	n,_ = strconv.Atoi(split_input[0])
	v,_ = strconv.Atoi(split_input[1])
	ans = "NO"
	for i = 1;i < n+1;i++{
		split_input = strings.Split(input[i]," ")
		t[i],_ = strconv.Atoi(split_input[0])
		p[i],_ = strconv.Atoi(split_input[1])
		for j = i; j > 0;j--{
			if i != j{
				if (p[j] - p[i]) / (t[j] - t[i]) > v{
					ans = "YES"
				}
			}
		}
	}
	fmt.Println(ans)
}
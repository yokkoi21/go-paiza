package main
import "fmt"
import "strings"
import "strconv"
import "os"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var arr [101]string
	var line []string
	var pass_score,n int
	var score,absense int
	i := 0
	for sc.Scan(){
		arr[i] = sc.Text()
		i++
	}
	line = strings.Split(arr[0]," ")
	n,_ = strconv.Atoi(line[0])
	pass_score,_ = strconv.Atoi(line[1])
	for i = 1; i <= n; i++{
		line = strings.Split(arr[i]," ")
		score,_ = strconv.Atoi(line[0])
		absense,_ = strconv.Atoi(line[1])
		absense_score := absense * 5
		if (score - absense_score) >= pass_score {
			fmt.Println(i)
		} else if pass_score == 0 {
			fmt.Println(i)
		}
	}
	
}
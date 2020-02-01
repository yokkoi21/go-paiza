package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [101]string
	var split_input []string
	var i int
	var balance,point int
	var fee int
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	n,_ := strconv.Atoi(split_input[0])
	m,_ := strconv.Atoi(split_input[1])
	balance = n
	for i = 1;i < m+1;i++{
		fee,_ = strconv.Atoi(input[i])
		if fee <= point{
			point -= fee
		}else{
			balance -= fee
			point += (fee / 10)
		}
		fmt.Printf("%d %d\n",balance,point)
	}
	
	
}
package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [191]string
	var split_input []string
	var n,m,i,j int
	var win,lose int
	
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ = strconv.Atoi(input[0])
	graph := make([][]string, n)
	for i:=0; i<n; i++{
		graph[i] = make([]string, n)
	}
	m = (n * (n-1)) / 2
	for i = 0; i < m;i++{
		split_input = strings.Split(input[i+1]," ")
		win,_ = strconv.Atoi(split_input[0])
		lose,_ = strconv.Atoi(split_input[1])
		graph[win-1][lose-1] = "W"
		graph[lose-1][win-1] = "L"
	}
	for i = 0; i < n;i++{
		graph[i][i] = "-"
	}
	for i = 0; i < n;i++{
		for j = 0;j < n;j++{
			if j == n - 1{
				fmt.Printf("%s\n",graph[i][j])
			}else{
				fmt.Printf("%s ",graph[i][j])
			}
		}
	}
}	
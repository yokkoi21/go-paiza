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
	var n,m,k int
	var i,j int
	var r_position_tmp int
	var r_position [51]int
	var ans_tmp [51]int

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	n,_ = strconv.Atoi(split_input[0])
	m,_ = strconv.Atoi(split_input[1])
	k,_ = strconv.Atoi(split_input[2])
	
	 
	for i = 1;i < h+1;i++{
		r_postion[i],_ = strconv.Atoi(input[i])
	}
	
	for i = 0;i < k;i++{
		for j = 1;j < m+1;j++{
			if r_position[]
		}
	}
	
	for i = 0;i < h;i++{
		for j = 0;j < w;j++{
			ans_tmp[i] += graph[i][j]
		}
		fmt.Println(ans_tmp[i])
		ans_tmp[i] = ""
	}
	
}
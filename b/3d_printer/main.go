package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [1000]string
	var split_input []string
	var i,j,k int
	var x,y,z int
	var count int
	
	

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	x,_ = strconv.Atoi(split_input[0])
	y,_ = strconv.Atoi(split_input[1])
	z,_ = strconv.Atoi(split_input[2])
	
	graph := make([][][]string, x)
	for i:=0; i<x; i++{
		graph[i] = make([][]string, y)
		for j:=0; j<y; j++{
			graph[i][j] = make([]string, z)
		}
	}
	ans_graph := make([][]string, y)
	for i:=0; i<x; i++{
		ans_graph[i] = make([]string, z)
		
	}	
	
	count = 0
	for k = 0;k < z;k++{
		for j = 0;j < y;j++{
			if input[count] == "--"{
				count++
			}else{
				split_input = strings.Split(input[count],"")
				for i = 0;i < x;i++{
					graph[i][j][k] = split_input[i]
				}
			}
			count++
		}
	}
	
	
}
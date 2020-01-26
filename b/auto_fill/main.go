package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [3]string
	var split_input1[]string
	var split_input2[]string
	var split_input3[]string
	var i int
	var h,w,a0_0,a0_1,a1_0,a1_1 int
	var h_diff,w_diff int

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	
	split_input1 = strings.Split(input[0]," ")
	split_input2 = strings.Split(input[1]," ")
	split_input3 = strings.Split(input[2]," ")

	h,_ = strconv.Atoi(split_input1[0])
	w,_ = strconv.Atoi(split_input1[1])
	a0_0,_ = strconv.Atoi(split_input2[0])
	a0_1,_ = strconv.Atoi(split_input2[1])
	a1_0,_ = strconv.Atoi(split_input3[0])
	a1_1,_ = strconv.Atoi(split_input3[1])

	graph := make([][]int, h)
	for i:=0; i<h; i++{
		graph[i] = make([]int, w)
	}	
	graph[0][0] = a0_0
	graph[0][1] = a0_1
	graph[1][0] = a1_0
	graph[1][1] = a1_1
	
	if graph[1][0] > graph[0][0] {
		h_diff =  graph[1][0] - graph[0][0]
	}else if graph[1][0] < graph[0][0] {
		h_diff =  graph[1][0] - graph[0][0] 
	}else {
		h_diff = 0
	}

	if graph[0][1] > graph[0][0] {
		w_diff =  graph[0][1] - graph[0][0]
	}else if graph[0][1] < graph[0][0] {
		w_diff =  graph[0][1] - graph[0][0]
	}else {
		w_diff = 0
	}
	
	for j := 0;j < h; j++{
		for k := 0;k < w; k++{
			if j == 0 && k == 0 {
			}else if j == 0 && k == 1{
				graph[j][k] = graph[0][1] 
			}else if j == 1 && k == 0{
				graph[j][k] = graph[1][0] 
			}else if j == 1 && k == 1{
				graph[j][k] = graph[1][1]
			}else if j == 0 {
				graph[j][k] = graph[0][k-1] + w_diff 
			}else if k == 0{
				graph[j][k] = graph[j-1][0] + h_diff 
			}else {
				graph[j][k] = graph[j][k-1] + (graph[j][1] - graph[j][0])
				graph[j][k] = graph[j-1][k] + (graph[1][k] - graph[0][k]) 
			}
		}	
	}
	for l := 0;l < h; l++{
		for m := 0;m < w; m++{
			if m < w - 1{
				fmt.Printf("%d ", graph[l][m])
			}else {
				fmt.Printf("%d\n", graph[l][m])
			}
			
		}	
	}

}
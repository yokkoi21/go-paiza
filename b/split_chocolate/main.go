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
	var h,w int
	var i,j int
	var h_sum [101]int
	var s_tmp int
	var check int
	var ans string

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	h,_ = strconv.Atoi(split_input[0])
	w,_ = strconv.Atoi(split_input[1])


	graph := make([][]int, h)
	for i:=0; i<h; i++{
		graph[i] = make([]int, w)
	}
	graph2 := make([][]string, h)
	for i:=0; i<h; i++{
		graph2[i] = make([]string, w)
	}
	 
	for i = 0;i < h;i++{
		split_input = strings.Split(input[i+1]," ")
		for j = 0;j < w;j++{
			graph[i][j],_ = strconv.Atoi(split_input[j])
			h_sum[i] += graph[i][j] 
		}
	}
	
	for i = 0;i < h;i++{
		for j = 0;j < w;j++{
			s_tmp += graph[i][j]
			if float64(s_tmp) == (float64(h_sum[i]) / float64(2)){
				check += 1
			}
			if float64(s_tmp) > (float64(h_sum[i]) / float64(2)){
				graph2[i][j] = "B"
			}else{
				graph2[i][j] = "A"
			}
			
		}
		s_tmp = 0
	}
	if check == h{
		fmt.Println("Yes")
		for i = 0;i < h;i++{
			for j = 0;j < w;j++{
				ans += graph2[i][j]
			}
			fmt.Println(ans)
			ans = ""
		}
	}else{
		fmt.Println("No")
	}
	
}
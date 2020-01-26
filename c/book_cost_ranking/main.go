package main
import "strings"
import "os"
import "strconv"
import "fmt"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [103]string
	var split_input1 []string
	var split_input2 []string
	var n,m,p,i,j int
	var o string
	var tmp_int int
	var tmp_str string
	cost := []int{}


	
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ = strconv.Atoi(input[0])
	
	split_input1 = strings.Split(input[1]," ")
    for i = 0;i<n;i++{
        cost = append(cost, 0)
    }
	m,_ = strconv.Atoi(input[2])
	for i = 0;i < m;i++{
		split_input2 = strings.Split(input[i+3]," ")
		o = split_input2[0]
		p,_ = strconv.Atoi(split_input2[1])
		for j = 0;j < n;j++{
			if o == split_input1[j]{
				cost[j] += p
			}
		}
	}
	for i = 0;i < n;i++{
		for j = n-1;j > i;j--{
			if cost[j] > cost[j-1]{
				tmp_int = cost[j]
				tmp_str = split_input1[j]
				cost[j] = cost[j-1]
				split_input1[j] = split_input1[j-1]
				cost[j -1] = tmp_int 
				split_input1[j -1] = tmp_str
			}
		}
	}
	
	for i = 0;i < n;i++{
		fmt.Println(split_input1[i])
	}
	
}
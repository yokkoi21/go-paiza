package main
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [102]string
	var split_input []string
	var word_list []string
	var i,n,j,k,k_num int
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	n,_ = strconv.Atoi(split_input[0])
	k_num,_ = strconv.Atoi(split_input[1])
	word_list = strings.Split(input[1]," ")
	for k = 0;k < n;k++{
		for i = 0;i < n;i++{
			for j = 0;j < 10;i++{
				if i == n - 1  || k == n - 1 {
					break
				}else if word_list[i][j] > word_list[i+1][j]{
					tmp := word_list[i]
					word_list[i] = word_list[i+1]
					word_list[i+1] = tmp 
					break
				}
			}
		}
	}
	
	tmp_num := n / k_num
	graph := make([][]string, tmp_num)
	for i:=0; i<tmp_num; i++{
    	graph[i] = make([]string, k_num)
	}
	for k = 0; k < len(word_list);k++{
		tmp_num2 := k % k_num
		for i = 0; i < tmp_num;i++{
			for j =0; j<k_num; i++{
    			graph[i][j] = word_list[tmp_num2+(i*k_num)]
			}
		}
	}
	
}
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
	var k,n int
	var i,j,i2,j2 int
	var tmp_num int
	var count int
	var tmp []string
	

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	n,_ = strconv.Atoi(split_input[0])
	k,_ = strconv.Atoi(split_input[1])
	
	graph := make([][]int, n)
	for i:=0; i<n; i++{
		graph[i] = make([]int, n)
	}
	graph2 := make([][]int, n)
	for i:=0; i<n; i++{
		graph2[i] = make([]int, n)
	}
	

	for i = 0;i < n;i++{
		tmp = strings.Split(input[i+1]," ")
		for j = 0;j < n;j++{
			graph[i][j],_ = strconv.Atoi(tmp[j])
			graph2[i][j],_ = strconv.Atoi(tmp[j])
		}
	}
	
	for i2 = 0;i2 < n / k;i2++{
		for j2 = 0;j2 < n / k;j2++{
			for i = 0+(i2*k);i < k+(i2*k); i++{
				for j = 0+(j2*k);j < k+(j2*k); j++{
					tmp_num += graph[i][j]
					count += 1
				}
			}
			graph2[i2][j2] = tmp_num / count
			tmp_num = 0
			count = 0
		}
	}
	
	
	for i = 0;i < n /k;i++{
		for j = 0;j < n /k;j++{
			if j != (n / k) - 1{
				fmt.Printf("%d ",graph2[i][j])
			}else{
				fmt.Printf("%d",graph2[i][j])
			}
			
		}
		fmt.Printf("\n")
	}
}
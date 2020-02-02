package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [11]string
	var split_input []string
	var n int
	var i,j int
	var sum,sum_max int
	var ans string
	var count int
	var count_check int

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	
	n,_ = strconv.Atoi(input[0])
	


	graph := make([][]int, n)
	for i:=0; i<n; i++{
		graph[i] = make([]int, n)
	}
	 
	
	for i = 1;i < n+1;i++{
		split_input = strings.Split(input[i]," ")
		for j = 0;j < n;j++{
			graph[i-1][j],_ = strconv.Atoi(split_input[j])
			sum += graph[i-1][j]
		}
		if sum > sum_max{
			sum_max = sum
		}
		sum = 0
	}
	
	for i = 0;i < n;i++{
		for j = 0;j < n;j++{
			sum += graph[i][j]
		}
		if sum < sum_max{
			for j = 0;j < n;j++{
				if graph[i][j] == 0{
					count += 1
					if sum_max - sum > n*n{
					}else{
						graph[i][j] = sum_max - sum
					}
				}
			}	
		}
		if count == 2{
			count_check = 1	
		}
		count = 0
		sum = 0
	}
	//fmt.Println(count_check)
	if count_check == 1{
		for  j = 0;j < n;j++{
			for i = 0;i < n;i++{
				sum += graph[i][j]
				//fmt.Println(sum)
			}
			if sum < sum_max{
				for i = 0;i < n;i++{
					if graph[i][j] == 0{
						graph[i][j] = sum_max - sum
					}
				}	
			}
			sum = 0			
		}
	}

	for i = 0;i < n;i++{
		for j = 0;j < n;j++{
			if j != n-1{
				ans = ans + strconv.Itoa(graph[i][j]) + " "
			}else {
				ans = ans + strconv.Itoa(graph[i][j])
			}
		}
		fmt.Println(ans)
		ans = ""
	}
	

}
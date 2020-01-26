package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var input [200]string
	var i,j int
	var m,n,k int
	var s,tmp,max int
	var supporter [101]int

	input = scan_input_str()
	
	m,n,k = split_and_atoi(input[0])

	supporter[0] = n
	for i = 1;i < k+1;i++{
		s,_ = strconv.Atoi(input[i])
		tmp = 0
		for j = 0;j < m+1;j++{
			if supporter[j] > 0{
				supporter[j] -= 1
				tmp += 1
			}
		}
		supporter[s] += tmp
	}
	
	max = 0
	for i = 1;i < m+1;i++{
		if supporter[i] >= max{
			max = supporter[i]
		}
	}
	for i = 1;i < m+1;i++{
		if supporter[i] == max{
			fmt.Println(i)
		}
	}

}

func scan_input_str() [200]string{
	var sc = bufio.NewScanner(os.Stdin)
	var input [200]string
	var i int
	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	
	return input
}

func split_and_atoi(input string) (int,int,int){
	var a,b,c int
	var split_input []string

	split_input = strings.Split(input," ")
	a,_ = strconv.Atoi(split_input[0])
	b,_ = strconv.Atoi(split_input[1])
	c,_ = strconv.Atoi(split_input[2])

	return a,b,c

}



package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [33]string
	var split_input []string
	var i,j int
	var c_n,s_n int
	var c_arr [100]int
	var s_slice []string
	var ans int

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	c_n,_ = strconv.Atoi(split_input[0])
	s_n,_ = strconv.Atoi(split_input[1])
	j = 0
	for i = 1;i < c_n + 1;i++{
		c_arr[j],_ = strconv.Atoi(input[i])
		j++
	}
	for i = c_n+1;i < s_n + c_n + 1;i++ {
		ans = 0
		s_slice = strings.Split(input[i]," ")
		for j = 0;j < c_n;j++{
			tmp,_ := strconv.Atoi(s_slice[j])
			ans += (c_arr[j] * tmp) / 100
		 } 
		 fmt.Println(ans)
	}
	
	
}
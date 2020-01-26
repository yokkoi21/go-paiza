package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [1001]string
	var split_input []string
	var i,j,n int
	var c_answer,s_answer []string
	var ans,count int

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ = strconv.Atoi(input[0])
	for i = 1;i < n + 1;i++ {
		split_input = strings.Split(input[i]," ")
		c_answer = strings.Split(split_input[0],"")
		s_answer = strings.Split(split_input[1],"")
		if split_input[0] == split_input[1]{
			ans += 2
		} else {
			if len(c_answer) == len(s_answer){
				count = 0
				for j = 0; j < len(c_answer); j++{
					if c_answer[j] != s_answer[j]{
						count += 1
					}
				}
				if count < 2{
					ans += 1
				}
			}
		}
	}
	fmt.Println(ans)
	
}
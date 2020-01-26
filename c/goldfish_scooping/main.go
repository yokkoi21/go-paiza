package main
import "os"
import "strconv"
import "strings"
import "fmt"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [2001]string
	var split_input []string
	var m,n,x,w,i int
	var tmp_x int
	var count int
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	m,_ = strconv.Atoi(split_input[0])
	n,_ = strconv.Atoi(split_input[1])
	x,_ = strconv.Atoi(split_input[2])
	tmp_x = x
	
	for i = 1;i < m+1;i++{
		w,_ = strconv.Atoi(input[i])
		//fmt.Println(input[i])
		if tmp_x - w > 0{
			count += 1
			tmp_x = tmp_x - w
		}else{
			n -= 1
			if n > 0{
				tmp_x = x
				i -= 1
			}else{
				break
			}
		}
	}
	
	fmt.Println(count)
}
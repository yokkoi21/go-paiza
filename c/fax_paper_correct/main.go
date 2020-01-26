package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [183]string
	var split_input []string
	var i int
	var n,h,m,p int
	var count int
	var tmp_h int
	var tmp_p int

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	
	m,_ = strconv.Atoi(input[0])
	n,_ = strconv.Atoi(input[1])
	split_input = strings.Split(input[2]," ")
	h,_ = strconv.Atoi(split_input[0])
	p,_ = strconv.Atoi(split_input[2])
	tmp_h = h
	tmp_p += p

	for i = 3;i < n+2;i++{
		split_input = strings.Split(input[i]," ")
		h,_ = strconv.Atoi(split_input[0])
		p,_ = strconv.Atoi(split_input[2])
		if tmp_h != h{
			if tmp_p % m == 0{
				count += tmp_p / m
			}else {
				count += (tmp_p / m) + 1
			}
			fmt.Println(count)
			tmp_p = 0
			tmp_h = h
		}else{
			tmp_p += p
		}
	}
	if tmp_p % m == 0{
		count += tmp_p / m
	}else {
		count += (tmp_p / m) + 1
	}
	
	fmt.Println(count)
	
}
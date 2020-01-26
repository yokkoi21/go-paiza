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
	var p1,p2,p3,p4,i int
	var p1_time,p2_time,p3_time,p4_time int
	var final1,final2 int
	var f1_time,f2_time int
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	p1,_ = strconv.Atoi(split_input[0])
	p2,_ = strconv.Atoi(split_input[1])
	split_input = strings.Split(input[1]," ")
	p3,_ = strconv.Atoi(split_input[0])
	p4,_ = strconv.Atoi(split_input[1])
	split_input = strings.Split(input[2]," ")
	p1_time,_ = strconv.Atoi(split_input[p1-1])
	p2_time,_ = strconv.Atoi(split_input[p2-1])
	p3_time,_ = strconv.Atoi(split_input[p3-1])
	p4_time,_ = strconv.Atoi(split_input[p4-1])
	if p1_time > p2_time{
		final1 = p2
	}else {
		final1 = p1
	}
	if p3_time > p4_time{
		final2 = p4
	}else {
		final2 = p3
	}
	split_input = strings.Split(input[3]," ")
	if final1 < final2{
		f1_time,_ = strconv.Atoi(split_input[0])
		f2_time,_ = strconv.Atoi(split_input[1])
	}else {
		f2_time,_ = strconv.Atoi(split_input[0])
		f1_time,_ = strconv.Atoi(split_input[1])
	}
	
	if f1_time < f2_time {
		fmt.Println(final1)
		fmt.Println(final2)
	}else {
		fmt.Println(final2)
		fmt.Println(final1)
	}
	
}
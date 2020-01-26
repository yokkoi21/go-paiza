package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input string
	var split_input []string
	var current_pattern []string
	var match_pattern_str string
	var current_pattern_str string
	var i,n int
	var tmp string
	var tmp_n int

	sc.Scan()
	input = sc.Text()
	split_input = strings.Split(input," ")
	current_pattern = strings.Split(split_input[2],"")
	match_pattern_str = split_input[1]
	current_pattern_str = split_input[2]
	//fmt.Println(current_pattern[4])
	n,_ = strconv.Atoi(split_input[0])
	//fmt.Printf("%c",split_input[1][0])
	for i = 0;i < n ;i++{
		if current_pattern_str == match_pattern_str{
			fmt.Println(i)
			break
		}else {
			tmp = current_pattern[0]
			// fmt.Println(tmp)
			tmp_n = n - 1
			current_pattern = current_pattern[1:]
			current_pattern = append(current_pattern,"?")
			//fmt.Println(current_pattern)
			current_pattern[tmp_n] = tmp
			current_pattern_str = strings.Join(current_pattern,"")
		}
		//fmt.Println(current_pattern[1])	
	}
	
	

}
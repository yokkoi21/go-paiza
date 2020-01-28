package main
import "strings"
import "os"
import "strconv"
import "fmt"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [1]string
	var split_input []string
	var page_list_str [10001]string
	var n,s,p int
	var i,j int
	var tmp string
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	n,_ = strconv.Atoi(split_input[0])
	s,_ = strconv.Atoi(split_input[1])
	p,_ = strconv.Atoi(split_input[2])
	
	
	for i = 0;i < n / s;i++{
		if n == 0{
			page_list_str[0] += "0"
		}else{		
			for j = (i * s);j < (s * (i + 1));j++{
				if j > n{
					break
				}else {
					if j != (s * (i + 1))-1{
						tmp = strconv.Itoa(j+1)
						page_list_str[i] += tmp + " "
					}else {
						tmp = strconv.Itoa(j+1)
						page_list_str[i] += tmp
					}
				}	
			}
		}
	}
	
	if page_list_str[p-1] == ""{
		fmt.Println("none")
	}else{
		fmt.Printf("%s\n",page_list_str[p-1])
	}
	
	
	
}
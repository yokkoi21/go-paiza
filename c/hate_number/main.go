package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [99999]string
	var i,j,m_int int
	var n string
	var r_list []string
	var no_print int
	var count int
	
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n = input[0]
	m_int,_ = strconv.Atoi(input[1])
	
	for i = 0;i < m_int; i++{
		no_print = 0
		r_list = strings.Split(input[i+2],"")
		for j = 0;j < len(r_list); j++{
			if r_list[j] == n {
				no_print = 1
			}
		}
		if no_print == 0 {
			fmt.Printf("%s\n",input[i+2])
			count += 1
		}
	}

	if count == 0{
		fmt.Printf("none\n")
	}
	fmt.Printf("\n")
	
}
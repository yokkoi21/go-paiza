package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [1020]string
	var split_input []string
	var m int
	var i int
	var octet1,octet2,octet3,octet4 int
	

	
	

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	m,_ = strconv.Atoi(split_input[0])
	
	for i = 0;i < m;i++{
		split_input = strings.Split(input[i+1],".")
		if split_input[0] != ""{
			octet1,_ = strconv.Atoi(split_input[0])
		}else{
			fmt.Println("False")
			continue
		}
		if split_input[1] != ""{
			octet2,_ = strconv.Atoi(split_input[1])
		}else{
			fmt.Println("False")
			continue
		}
		if split_input[2] != ""{
			octet3,_ = strconv.Atoi(split_input[2])
		}else{
			fmt.Println("False")
			continue
		}
		if split_input[3] != ""{
			octet4,_ = strconv.Atoi(split_input[3])
		}else{
			fmt.Println("False")
			continue
		}

		if len(split_input) != 4{
			fmt.Println("False")
			continue
		}
		

		if octet1 <= 255 && octet1 >= 0 && octet2 <= 255 && octet2 >= 0 && octet3 <= 255 && octet3 >= 0 && octet4 <= 255 && octet4 >= 0{
			fmt.Println("True")
			continue
		}else{
			fmt.Println("False")
			continue
		}
	}
	
}
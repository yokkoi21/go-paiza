package main
import "fmt"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [31]string
	var n,r_flag,s_flag,p_flag,i int
	
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ = strconv.Atoi(input[0])
	
	for i = 1;i < n+1; i++{
		if input[i] == "rock"{
			r_flag  = 1
		}else if input[i] == "scissors"{
			s_flag  = 1
		}else if input[i] == "paper"{
			p_flag  = 1
		}
	}

	if r_flag == 1 && s_flag == 1 && p_flag == 1{
		fmt.Println("draw")
	}else if r_flag == 1 && s_flag == 1 && p_flag == 0{
		fmt.Println("rock")
	}else if r_flag == 0 && s_flag == 1 && p_flag == 1{
		fmt.Println("scissors")
	}else if r_flag == 1 && s_flag == 0 && p_flag == 1{
		fmt.Println("paper")
	}else {
		fmt.Println("draw")
	}
	
	
}
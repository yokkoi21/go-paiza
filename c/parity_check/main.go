package main
import "strings"
import "os"
import "strconv"
import "fmt"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [51]string
	var split_input []string
	var n,i int
	var b1,b2,b3,b4 int
	var b1_tmp,b2_tmp,b3_tmp,b4_tmp int
	var ans string
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ = strconv.Atoi(input[0])
	
	for i = 1;i < n + 1;i++{
		split_input = strings.Split(input[i],"")
		b1,_ = strconv.Atoi(split_input[0])
		b2,_ = strconv.Atoi(split_input[1])
		b3,_ = strconv.Atoi(split_input[2])
		b4,_ = strconv.Atoi(split_input[3])
		if b1 == 1 && b1_tmp == 1{
			b1_tmp = 0
		}else{
			b1_tmp += b1
		}
		if b2 == 1 && b2_tmp == 1{
			b2_tmp = 0
		}else{
			b2_tmp += b2
		}
		if b3 == 1 && b3_tmp == 1{
			b3_tmp = 0
		}else{
			b3_tmp += b3
		}
		if b4 == 1 && b4_tmp == 1{
			b4_tmp = 0
		}else{
			b4_tmp += b4
		}

	}
	
	ans += strconv.Itoa(b1_tmp)
	ans += strconv.Itoa(b2_tmp)
	ans += strconv.Itoa(b3_tmp)
	ans += strconv.Itoa(b4_tmp)
	
	fmt.Println(ans)
	
	
}
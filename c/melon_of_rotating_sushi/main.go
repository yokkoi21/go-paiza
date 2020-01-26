package main
import "os"
import "strconv"
import "fmt"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [2001]string
	var t,i int
	var count int
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	t,_ = strconv.Atoi(input[0])
	
	for i = 1;i < t+1;i++{
		if input[i] == "melon"{
			count += 1
			i += 10
		}
	}
	
	fmt.Println(count)
}
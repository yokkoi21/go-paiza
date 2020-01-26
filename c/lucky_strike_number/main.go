package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [102]string
	var strile_number []string
	var my_number []string
	var i,n,j,k int
	var ans int
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	strile_number = strings.Split(input[0]," ")
	n,_ = strconv.Atoi(input[1])
	
	for i = 2;i < n +2;i++{
		ans = 0
		my_number = strings.Split(input[i]," ")
		for j = 0; j < 6;j++{
			for k = 0;k < 6;k++{
				if strile_number[j] == my_number[k]{
					ans += 1
				}
			}
		}
		fmt.Println(ans)
	}

	
}
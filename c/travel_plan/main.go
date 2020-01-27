package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [60]string
	var split_input []string
	var m,n,d,i,j,k int
	var tmp,tmp_p int
	var min float64
	var s_arr,e_arr [31]int
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	m,_ = strconv.Atoi(split_input[0])
	n,_ = strconv.Atoi(split_input[1])
	min = 100
	for i = 1;i < (m + 1)- (n - 1); i++{
		split_input = strings.Split(input[i]," ")
		d,_ = strconv.Atoi(split_input[0])
		for j = i;j < i+n;j++{
			split_input = strings.Split(input[j]," ")
			tmp_p,_ = strconv.Atoi(split_input[1])
			tmp += tmp_p
		}
		//fmt.Println(tmp)
		if float64(tmp) / float64(n) < min{
			min = float64(tmp) / float64(n)
		}
		tmp = 0
	}

	for i = 1;i < (m + 1)- (n - 1); i++{
		split_input = strings.Split(input[i]," ")
		d,_ = strconv.Atoi(split_input[0])
		for j = i;j < i+n;j++{
			split_input = strings.Split(input[j]," ")
			tmp_p,_ = strconv.Atoi(split_input[1])
			tmp += tmp_p
		}
		//fmt.Println(tmp)
		if float64(tmp) / float64(n) == min{
			s_arr[k] = d
			e_arr[k] = d + n - 1
			k++
		}
		tmp = 0
	}
	
	fmt.Printf("%d %d",s_arr[0],e_arr[0])
	
	
}
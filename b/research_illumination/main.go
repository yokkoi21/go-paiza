package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [1003]string
	var split_input []string
	var n,m,q int
	var right_num [1000]int
	var i,j int
	var start,end int
	var amount,heikin,count int
	
	i = 0
	j = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	
	split_input = strings.Split(input[0]," ")
	n,_ = strconv.Atoi(split_input[0])
	m,_ = strconv.Atoi(split_input[1])
	split_input = strings.Split(input[1]," ")
	for i = 0;i < n;i++{
		right_num[i],_ = strconv.Atoi(split_input[i])
	}
	q,_ = strconv.Atoi(input[2])
	for i = 3;i < q+3;i++{
		count = 0
		amount = 0
		heikin = 0
		split_input = strings.Split(input[i]," ")
		start,_ = strconv.Atoi(split_input[0])
		end,_ = strconv.Atoi(split_input[1])
		for j = start - 1; j < end; j++{
			amount += right_num[j]
			count += 1 
		}
		heikin = amount / count
		
		for m > heikin {
			amount = 0			
			for j = start - 1; j < end; j++{
				right_num[j] += 1  
				amount += right_num[j]
			}
			heikin = amount / count
		}
	}

	for i = 0;i < n;i++{
		if i == n - 1{
			fmt.Printf("%d\n",right_num[i])
		}else {
			fmt.Printf("%d ",right_num[i])
		}
		
	}

}
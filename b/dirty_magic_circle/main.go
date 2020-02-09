package main
import "fmt"
//import "strings"
import "os"
import "strconv"
import "bufio"
import "sort"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [20]string
	var s,n int
	var i,j int
	var count int
	var product []int
	var pattern []int
	var tmp int
	var tmp_value int
	

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	s,_ = strconv.Atoi(input[0])
	n,_ = strconv.Atoi(input[1])
	
	 
	
	for i = 0;i < n;i++{
		tmp,_ = strconv.Atoi(input[i+2])
		product = append(product, tmp)
	}
	
	sort.Sort(sort.IntSlice(product))
	//fmt.Println(product)

	for i = n-1;i >= 0;i--{
		tmp_value = 0
		tmp_value += product[i]
		if tmp_value >= s{
			count += 1
			pattern = append(pattern, tmp_value)
			continue
		}else{
			for j = i-1;j >= 0;j--{
				tmp_value += product[j]
				if tmp_value >= s{
					count += 1
					pattern = append(pattern, tmp_value)
					tmp_value -= product[j]
				}
			}
			
			// for j = 0;j < i;j++{
			// 	tmp_value -= product[j]
			// 	if tmp_value < s{
			// 		count -= 1
			// 		break
			// 	}
			// }
		}
	}
	
	fmt.Println(count)
	//fmt.Println(pattern)

}
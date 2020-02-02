package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [202]string
	var split_input []string
	var i,j int
	var n,m int
	var n_material_name [101]string
	var n_material_num [101]int
	var o_material_name [10001]string
	var o_material_num [10001]int
	var shortage_check int
	var count int
	var material_check int


	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ = strconv.Atoi(input[0])

	for i = 1;i < n + 1;i++{
		split_input = strings.Split(input[i]," ")
		n_material_name[i] = split_input[0]
		n_material_num[i],_ = strconv.Atoi(split_input[1])
	}
	m,_ = strconv.Atoi(input[n+1])
	for i =  n + 2;i < n + m + 2;i++{
		split_input = strings.Split(input[i]," ")
		o_material_name[i-n-1] = split_input[0]
		o_material_num[i-n-1],_ = strconv.Atoi(split_input[1])
	}
	// fmt.Println(n_material_name[1])
	// fmt.Println(o_material_name[1])
	for {
		for i = 1;i < n + 1;i++{
			for j = 1;j < m + 1;j++{
				if o_material_name[j] == n_material_name[i] {
					if o_material_num[j] - n_material_num[i] < 0{
						shortage_check = 1
					}else{
						o_material_num[j] -= n_material_num[i]
						material_check += 1
					}
				}
			}
			if  material_check == 0{
				shortage_check = 1
			} 
		}
		//fmt.Println(material_check)
		if material_check == n{
			material_check = 0
		}else{
			shortage_check = 1
		}
		if shortage_check == 1{
			break
		}else{
			count += 1
		}
	}
	fmt.Println(count)
}
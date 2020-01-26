package main
import "strings"
import "os"
import "strconv"
import "fmt"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [2]string
	var split_input1 []string
	var split_input2 []string
	var y,m,d,a,b,i,j,k int
	var count int



	
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input1 = strings.Split(input[0]," ")
	y,_ = strconv.Atoi(split_input1[0])
	m,_ = strconv.Atoi(split_input1[1])
	d,_ = strconv.Atoi(split_input1[2])
	split_input2 = strings.Split(input[1]," ")
	a,_ = strconv.Atoi(split_input2[0])
	b,_ = strconv.Atoi(split_input2[1])

	for i = 0;i <4;i++{
		for j = m; j < 14;j++{
			for k = d; k < 16;k++{
				if (j % 2) == 1 && k > 13{
					break
				}
				if (y + i) % 4 == 1{
					if j == a && k == b{
						break
					}
				}
				count += 1
			}
			if j == a && k == b{
				break
			}
			d = 1
		}
		if j == a && k == b{
			break
		}
		m = 1
	}
	
	fmt.Println(count)
}
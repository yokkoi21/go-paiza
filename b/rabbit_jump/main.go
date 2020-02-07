package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [101]string
	var split_input []string
	var n,m,k int
	var i,j,l int
	var r_position [501]int
	var r_position_tmp [501]int

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	n,_ = strconv.Atoi(split_input[0])
	m,_ = strconv.Atoi(split_input[1])
	k,_ = strconv.Atoi(split_input[2])
	
	// for i = 1;i < 501;i++{
	// 	r_position_tmp[r_position[i]] = 0
	// }
	for i = 1;i < m+1;i++{
		r_position[i],_ = strconv.Atoi(input[i])
		r_position_tmp[r_position[i]] = 1
	}
	// for i = 1;i < m+1;i++{
	// 	if r_position[i] != 0{
	// 		fmt.Printf("%d",r_position[i])
	// 	}	
	// }
	// fmt.Printf("\n")

	// for i = 1;i < 10;i++{
	// 	fmt.Printf("%d",r_position_tmp[i])	
	// }
	
	for i = 0;i < k;i++{
		for j = 1;j < m+1;j++{
			for l = 1;l < n;l++{
				if r_position[j]+l <= n{
					if r_position_tmp[r_position[j]+l] == 0{
						r_position[j] = r_position[j]+l
						r_position_tmp[r_position[j]] = 1
						r_position_tmp[r_position[j]-l] = 0
						break
					}else{
						continue
					}
				}else{
					if r_position_tmp[r_position[j]+l-n] == 0{
						r_position[j] = r_position[j]+l-n
						r_position_tmp[r_position[j]] = 1
						r_position_tmp[r_position[j]-l+n] = 0
						break
					}else{
						continue
					}
				}
			}
		}
	}
	//fmt.Printf("\n")
	for i = 0;i < n;i++{
		if r_position[i] != 0{
			fmt.Println(r_position[i])
		}	
	}
	// for i = 1;i < m+1;i++{
	// 	if r_position[i] != 0{
	// 		fmt.Printf("%d",r_position[i])
	// 	}	
	// }
	// fmt.Printf("\n")

	// for i = 1;i < 10;i++{
	// 	fmt.Printf("%d",r_position_tmp[i])	
	// }
	
}
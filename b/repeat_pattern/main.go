package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [20]string
	var k,n int
	var i,j,l,i2,j2 int
	var tmp_num int
	var tmp_num2 int
	var tmp []string
	

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	k,_ = strconv.Atoi(input[0])
	n,_ = strconv.Atoi(input[1])
	tmp_num = n
	for i = 0;i < k;i++{
		tmp_num = tmp_num * tmp_num
	}
	graph := make([][]string, tmp_num)
	for i:=0; i<tmp_num; i++{
		graph[i] = make([]string, tmp_num)
	}
	graph2 := make([][]string, tmp_num)
	for i:=0; i<tmp_num; i++{
		graph2[i] = make([]string, tmp_num)
	}

	for i = 0;i < n;i++{
		tmp = strings.Split(input[i+2],"")
		for j = 0;j < n;j++{
			graph[i][j] = tmp[j]
			graph2[i][j] = tmp[j]
		}
	}
	
	tmp_num = n
	tmp_num2 = n
	for l = 1;l < k+1;l++{
		for i = 0;i < tmp_num;i++{
			for j = 0;j < tmp_num;j++{
				if graph[i][j] == "#"{
					for i2 = i*tmp_num;i2 < tmp_num*tmp_num;i2++{
						for j2 = j*tmp_num2;j2 < tmp_num2*tmp_num2;j2++{
							graph2[i2][j2] = graph[i2][j2]
							graph[i2][j2] = graph2[i2-(i*tmp_num)][j2-(j*tmp_num)]
							
							// for t := 0;t < 16;t++{
							// 	for y := 0;y < 16;y++{
							// 		fmt.Printf("%s",graph[t][y])
							// 	}
							// 	fmt.Printf("\n")
							// }
						}
					}
				}else if graph[i][j] == "."{
					for i2 = i*tmp_num;i2 < tmp_num*tmp_num;i2++{
						for j2 = j*tmp_num2;j2 < tmp_num2*tmp_num2;j2++{
							graph[i2][j2] = "."
							// for t := 0;t < 16;t++{
							// 	for y := 0;y < 16;y++{
							// 		fmt.Printf("%s",graph[t][y])
							// 	}
							// 	fmt.Printf("\n")
							// }	
						}
					}
				}else {
					for i2 = i*tmp_num;i2 < tmp_num*tmp_num;i2++{
						for j2 = j*tmp_num2;j2 < tmp_num2*tmp_num2;j2++{
							graph2[i2][j2] = graph[i2][j2]
							graph[i2][j2] = graph2[i2-(i*tmp_num)][j2-(j*tmp_num)]
							
							// for t := 0;t < 16;t++{
							// 	for y := 0;y < 16;y++{
							// 		fmt.Printf("%s",graph[t][y])
							// 	}
							// 	fmt.Printf("\n")
							// }
						}
					}
				}

			}
		}
		tmp_num = tmp_num * tmp_num
		tmp_num2 = tmp_num2 * tmp_num2
	}
	
	for i = 0;i < tmp_num;i++{
		for j = 0;j < tmp_num;j++{
			fmt.Printf("%s",graph[i][j])
		}
		fmt.Printf("\n")
	}
}
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
	var h,w int
	var h_s,w_s int
	var i,j int
	var direction string
	var ans_tmp [51]string

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	h,_ = strconv.Atoi(split_input[0])
	w,_ = strconv.Atoi(split_input[1])
	split_input = strings.Split(input[1]," ")
	h_s,_ = strconv.Atoi(split_input[0])
	w_s,_ = strconv.Atoi(split_input[1])


	graph := make([][]string, h)
	for i:=0; i<h; i++{
		graph[i] = make([]string, w)
	}
	
	 
	for i = 0;i < h;i++{
		split_input = strings.Split(input[i+2],"")
		for j = 0;j < w;j++{
			graph[i][j] = split_input[j]
			
		}
	}
	// for i = 0;i < h;i++{
	// 	for j = 0;j < w;j++{
	// 		ans_tmp[i] += graph[i][j]
	// 	}
	// 	fmt.Println(ans_tmp[i])
	// 	ans_tmp[i] = ""
	// }
	direction = "N"
	h_s -= 1
	w_s -= 1
	for i = 0;i < 2000;i++{
		if h_s < h && w_s < w && h_s >= 0 && w_s >= 0{ 
			if graph[h_s][w_s] == "."{
				graph[h_s][w_s] = "*"
					if direction == "N"{
						w_s += 1
						direction = "E"
						continue
					}else if direction == "E" {
						h_s += 1
						direction = "S"
						continue
					}else if direction == "S" {
						w_s -= 1
						direction = "W"
						continue
					}else if direction == "W" {
						h_s -= 1
						direction = "N"
						continue
					}
			}else if graph[h_s][w_s] == "*"{
				graph[h_s][w_s] = "."
				if direction == "N"{	
					w_s -= 1
					direction = "W"
					continue
				}else if direction == "W" {
					h_s += 1
					direction = "S"
					continue
				}else if direction == "S" {
					w_s += 1
					direction = "E"
					continue
				}else if direction == "E" {
					h_s -= 1
					direction = "N"
					continue
				}
			}
		}else {
				break
		}
	}
	
	for i = 0;i < h;i++{
		for j = 0;j < w;j++{
			ans_tmp[i] += graph[i][j]
		}
		fmt.Println(ans_tmp[i])
		ans_tmp[i] = ""
	}
	
}
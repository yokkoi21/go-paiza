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
	var w,h int
	var i,j int
	var header [101]string
	var string_max [101]int
	var h_bar string
	

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	w,_ = strconv.Atoi(input[0])
	split_input = strings.Split(input[1]," ")
	for i = 0;i < w;i++{
		header[i] = split_input[i]
		if string_max[i] <= len(header[i]){
			string_max[i] = len(header[i])
			//fmt.Println(string_max[i])
		}
	}

	h,_ = strconv.Atoi(input[2])
	graph := make([][]string, h)
	for i:=0; i<h; i++{
		graph[i] = make([]string, w)
	}
	
	for i = 0;i < h;i++{
		split_input = strings.Split(input[i+3]," ")
		for j = 0;j < w;j++{
			graph[i][j] = split_input[j]
			if string_max[j] <= len(graph[i][j]){
				string_max[j] = len(graph[i][j])
				//fmt.Println(string_max[j])
			}
		}
	}

	
	for i = 0;i < w;i++{
		fmt.Printf("| ")
		if string_max[i] > len(header[i]){
			for {
				header[i] += " "
				if string_max[i] <= len(header[i]){
					break
				}
			}
		}
		fmt.Printf("%s ",header[i])
	}
	fmt.Printf("|\n")

	for i = 0;i < w;i++{
		fmt.Printf("|-")
		if string_max[i] > len(h_bar){
			for {
				h_bar += "-"
				if string_max[i] <= len(h_bar){
					break
				}
			}
		}
		fmt.Printf("%s-",h_bar)
		h_bar = ""
	}
	fmt.Printf("|\n")

	for i = 0;i < h;i++{
		for j = 0;j < w;j++{
			fmt.Printf("| ")
			if string_max[j] > len(graph[i][j]){
				for {
					graph[i][j] += " "
					if string_max[j] <= len(graph[i][j]){
						break
					}
				}
			}
			fmt.Printf("%s ",graph[i][j])
		}
		fmt.Printf("|\n")
	}
}
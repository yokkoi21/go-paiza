package main
import "strings"
import "os"
import "strconv"
import "fmt"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [101]string
	var split_input []string
	var g [101]int
	var s [101]int
	var b [101]int
	var n,i,j int
	var tmp_g,tmp_s,tmp_b int
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ = strconv.Atoi(input[0])
	
	
	for i = 1;i < n+1;i++{
		split_input = strings.Split(input[i]," ")
		g[i],_ = strconv.Atoi(split_input[0])
		s[i],_ = strconv.Atoi(split_input[1])
		b[i],_ = strconv.Atoi(split_input[2])
		
	}
	for i = 1;i < n+1;i++{
		for j = n;j > i;j--{
			if g[j] > g[j-1]{
				tmp_g = g[j]
				tmp_s = s[j]
				tmp_b = b[j]
				g[j] = g[j-1]
				s[j] = s[j-1]
				b[j] = b[j-1]
				g[j -1] = tmp_g
				s[j -1] = tmp_s
				b[j -1] = tmp_b 
			}else if g[j] == g[j -1]{
				if s[j] > s[j-1]{
					tmp_g = g[j]
					tmp_s = s[j]
					tmp_b = b[j]
					g[j] = g[j-1]
					s[j] = s[j-1]
					b[j] = b[j-1]
					g[j -1] = tmp_g
					s[j -1] = tmp_s
					b[j -1] = tmp_b 
				}else if s[j] == s[j -1]{
					if b[j] > b[j-1]{
						tmp_g = g[j]
						tmp_s = s[j]
						tmp_b = b[j]
						g[j] = g[j-1]
						s[j] = s[j-1]
						b[j] = b[j-1]
						g[j -1] = tmp_g
						s[j -1] = tmp_s
						b[j -1] = tmp_b 
					}
				}
			}
		}
	}
	
	for i = 1;i < n+1;i++{
		fmt.Printf("%d %d %d\n",g[i],s[i],b[i])
	}
	
}
package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var arr [1001]string
	var split_arr []string
	var i int
	for sc.Scan(){
		arr[i] = sc.Text()
		i++
	}
	var n int 
	n,_ = strconv.Atoi(arr[0])
	for i = 1;i < n+1; i++{
		if strings.HasSuffix(arr[i], "s") || strings.HasSuffix(arr[i], "sh") || strings.HasSuffix(arr[i], "ch") || strings.HasSuffix(arr[i], "o") || strings.HasSuffix(arr[i], "x"){
			arr[i] += "es"
		}else if strings.HasSuffix(arr[i], "f"){
			split_arr = strings.Split(arr[i],"")
			tmp := len(arr[i]) - 1 
			split_arr[tmp] = ""
			arr[i] = strings.Join(split_arr,"") 
			arr[i] += "ves"
		}else if strings.HasSuffix(arr[i], "fe"){
			split_arr = strings.Split(arr[i],"")
			tmp1 := len(arr[i]) - 1
			tmp2 := len(arr[i]) - 2
			split_arr[tmp1] = ""
			split_arr[tmp2] = ""
			arr[i] = strings.Join(split_arr,"") 
			arr[i] += "ves"
		}else if strings.HasSuffix(arr[i], "y"){
			if strings.HasSuffix(arr[i], "ay"){
				arr[i] += "s"
			}else if strings.HasSuffix(arr[i], "iy"){
				arr[i] += "s"
			}else if strings.HasSuffix(arr[i], "uy"){
				arr[i] += "s"
			}else if strings.HasSuffix(arr[i], "ey"){
				arr[i] += "s"
			}else if strings.HasSuffix(arr[i], "oy"){
				arr[i] += "s"
			}else {
				split_arr = strings.Split(arr[i],"")
				tmp := len(arr[i]) - 1 
				split_arr[tmp] = ""
				arr[i] = strings.Join(split_arr,"") 
				arr[i] += "ies"
			}
			
		}else {
			arr[i] += "s"
		}
		fmt.Println(arr[i])
	}
}
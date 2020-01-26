package main
import "fmt"
import "strconv"
import "os"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var arr1 [100]int
	var arr2 [100]int
	i := 0
	l := 0
	for sc.Scan(){
		arr1[i],_ = strconv.Atoi(sc.Text())
		i += 1
	}

	for i = 0; i < len(arr1); i++{
		for l = 1; l < arr1[i]; l++{
			if arr1[i] % l == 0 {
				arr2[i] += l
			} 
		}
	}

	for i = 1; i < len(arr2); i++{
		if arr1[i] == arr2[i] && arr1[i] != 0 {
			fmt.Println("perfect")
		}else if arr1[i] - 1 == arr2[i] && arr1[i] != 0{
			fmt.Println("nearly")
		}else if arr1[i] == 0  {
		}else {
		    
			fmt.Println("neither")
		}
	}
}
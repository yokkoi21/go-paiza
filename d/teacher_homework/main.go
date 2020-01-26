 package main
import "fmt"
import "strings"
import "os"
import "bufio"
import "strconv"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var arr string
	var split_arr []string
	var a,b,c int
	var ans int
	sc.Scan()
	arr = sc.Text()
	split_arr = strings.Split(arr," ")
	if split_arr[0] == "x"{
		b,_ = strconv.Atoi(split_arr[2])
		c,_ = strconv.Atoi(split_arr[4])
		if split_arr[1] == "-"{
			ans = b + c
		}else if split_arr[1] == "+"{
			ans = c - b
		}
	}else if split_arr[2] == "x"{
		a,_ = strconv.Atoi(split_arr[0])
		c,_ = strconv.Atoi(split_arr[4])
		if split_arr[1] == "-"{
			ans = a - c
		}else if split_arr[1] == "+"{
			ans = c - a
		}
	}else if split_arr[4] == "x"{
		a,_ = strconv.Atoi(split_arr[0])
		b,_ = strconv.Atoi(split_arr[2])
		if split_arr[1] == "-"{
			ans = a - b
		}else if split_arr[1] == "+"{
			ans = a + b
		}
	}

	fmt.Println(ans)

}
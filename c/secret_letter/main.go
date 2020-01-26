package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [191]string
	var n,i int
	var cipher_text []string
	var ans string
	
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ = strconv.Atoi(input[0])
	cipher_text = strings.Split(input[1],"")

	for i = 0;i < len(cipher_text);i++{
		if i % 2 == 0{
			if (int(cipher_text[i][0]) - n) <  65{
				ans += string(int(cipher_text[i][0]) - n + 26)
			}else {
				ans += string(int(cipher_text[i][0]) - n)
			}
		}else {
			if (int(cipher_text[i][0]) + n) >  90{
				ans += string(int(cipher_text[i][0]) + n - 26)
			}else {
				ans += string(int(cipher_text[i][0]) + n)
			}
		}
	}

	fmt.Println(ans)
	
}	
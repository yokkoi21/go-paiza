package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input string
	var i,j int
	var split_input []string
	var m,n int
	var count int
	

	sc.Scan()
	input = sc.Text()

	split_input = strings.Split(input," ")
	m,_ = strconv.Atoi(split_input[0])
	n,_ = strconv.Atoi(split_input[1])
	

	for i = 1;i < m+1;i++{
		for j = 1;j < n+1;j++{
			if i % 3 == 0 && j % 4 == 0{
				if ((i * i) + (j * j)) % 25 == 0  {
					count += 1
				}
			}else if i % 4 == 0 && j % 3 == 0{
				if ((i * i) + (j * j)) % 25 == 0 {
					count += 1
				}
			}else if i % 5 == 0 && j % 12 == 0{
				if ((i * i) + (j * j)) % 169 == 0  {
					count += 1
				}
			}else if i % 12 == 0 && j % 5 == 0{
				if ((i * i) + (j * j)) % 169 == 0 {
					count += 1
				}
			}else if i % 24 == 0 && j % 7 == 0{
				if ((i * i) + (j * j)) % 625 == 0  {
					count += 1
				}
			}else if i % 7 == 0 && j % 24 == 0{
				if ((i * i) + (j * j)) % 625 == 0 {
					count += 1
				}
			}else  if i % 15 == 0 && j % 8 == 0{
				if ((i * i) + (j * j)) % 289 == 0  {
					count += 1
				}
			}else if i % 8 == 0 && j % 15 == 0{
				if ((i * i) + (j * j)) % 289 == 0 {
					count += 1
				}
			}else  if i % 21 == 0 && j % 20 == 0{
				if ((i * i) + (j * j)) % 841 == 0  {
					count += 1
				}
			}else if i % 20 == 0 && j % 21 == 0{
				if ((i * i) + (j * j)) % 841 == 0 {
					count += 1
				}
			}
		}
	}

	fmt.Printf("%d\n",count)
	
	
}	
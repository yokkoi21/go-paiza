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
	var m,n int
	var i int
	var container_max int
	var container_count int
	var container int
	var amari int
	var amari_min int
	var ans int
	

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	m,_ = strconv.Atoi(split_input[0])
	n,_ = strconv.Atoi(split_input[1])


	amari_min = 1000
	for i = 1;i < m+1;i++{
		container,_ = strconv.Atoi(input[i])
		container_count = container
		amari = n % container
		//fmt.Println(amari)
		if amari <= amari_min{
			amari_min = amari
		}
	}
	for i = 1;i < m+1;i++{
		container,_ = strconv.Atoi(input[i])
		container_count = container
		amari = n % container
		//fmt.Println(amari)
		if amari == amari_min{
			if container_count > container_max{
				container_max = container_count
				//fmt.Println(container_count)
				ans = i
			}
		}
	}

	fmt.Println(ans)
}
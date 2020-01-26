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
	var ans int
	var i int
	var sum int
	var l_score_sum,s_score_sum int
	for sc.Scan(){
		arr[i] = sc.Text()
		i++
	}
	var n int 
	n,_ = strconv.Atoi(arr[0])
	for i = 1;i < n+1; i++{
		split_arr = strings.Split(arr[i]," ")
		sum = 0
		e_score,_ :=  strconv.Atoi(split_arr[1])
		m_score,_ :=  strconv.Atoi(split_arr[2])
		s_score,_ :=  strconv.Atoi(split_arr[3])
		j_score,_ :=  strconv.Atoi(split_arr[4])
		g_score,_ :=  strconv.Atoi(split_arr[5])
		sum = e_score + m_score + s_score + j_score + g_score
		if sum >= 350{
			if split_arr[0] == "l"{
				l_score_sum = j_score + g_score
				if l_score_sum >= 160{
					ans += 1
				}
			}else if split_arr[0] == "s"{
				s_score_sum = m_score + s_score
				if s_score_sum >= 160{
					ans += 1
				}
			}
		}
	}

	fmt.Println(ans)

}
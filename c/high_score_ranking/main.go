package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"
import "math"
import "sort"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [1010]string
	var split_input []string
	var n,m,k int
	var i,j int
	var tmp_item_score_f float64
	var tmp_item_score []string
	var tmp_item_score1 int
	var tmp_item_score2 int
	var item_score [110]float64
	var total_tmp_f float64
	var total_tmp_i int
	var score_total []int

	
	

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	n,_ = strconv.Atoi(split_input[0])
	m,_ = strconv.Atoi(split_input[1])
	k,_ = strconv.Atoi(split_input[2])

	graph := make([][]int, m)
	for i:=0; i<m; i++{
		graph[i] = make([]int, n)
	}
	
	split_input = strings.Split(input[1]," ")
	for i = 0;i < n;i++{
		tmp_item_score = strings.Split(split_input[i],".")
		tmp_item_score1,_ = strconv.Atoi(tmp_item_score[0])
		if len(tmp_item_score) == 2{
			tmp_item_score2,_ = strconv.Atoi(tmp_item_score[1])
		}else{
			tmp_item_score2 = 0
		}
		//fmt.Printf("%d\n",tmp_item_score2)
		tmp_item_score_f = float64(tmp_item_score1) + ((float64(tmp_item_score2) / 10.0))
		item_score[i] = tmp_item_score_f
	}
	
	for i = 0;i < m;i++{
		split_input = strings.Split(input[i+2]," ")
		for j = 0;j < n;j++{
			graph[i][j],_ = strconv.Atoi(split_input[j])
			total_tmp_f += float64(graph[i][j]) * item_score[j]
			//fmt.Println(total_tmp_f)
		}
		//fmt.Printf("\n")
		total_tmp_i = int(math.Round(total_tmp_f*1) / 1)
		score_total = append(score_total,total_tmp_i)
		total_tmp_f = 0
	}

	sort.Slice(score_total, func(i,j int) bool {
        return score_total[i] > score_total[j]
    })

	for i = 0;i < k;i++{
		fmt.Println(score_total[i])
	}
}
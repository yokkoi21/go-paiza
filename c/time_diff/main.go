package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [102]string
	var split_input []string
	var i int
	var n int
	var city [101]string
	var diff_time [101]int
	var base_time string
	var base_city string
	var tmp_h,base_h int
	var base_m string
	var s_time_h int

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	n,_ = strconv.Atoi(input[0])

	split_input = strings.Split(input[n+1]," ")
	base_city = split_input[0]
	base_time = split_input[1]
	split_input = strings.Split(base_time,":")
	base_h,_ = strconv.Atoi(split_input[0])
	base_m = split_input[1]

	for i = 1;i < n+1;i++{
		split_input = strings.Split(input[i]," ")
		city[i] = split_input[0]
		diff_time[i],_ = strconv.Atoi(split_input[1])
		if base_city == city[i]{
			if base_h - diff_time[i] >= 0 && base_h - diff_time[i] < 24{
				s_time_h = base_h - diff_time[i]
			}else if base_h - diff_time[i] >= 24{
				s_time_h = base_h - diff_time[i] - 24
			}else if base_h - diff_time[i] < 0{
				s_time_h = base_h - diff_time[i] + 24
			} 
			
		}
		
	}
	//fmt.Printf("%d",s_time_h)
	for i = 1;i < n+1;i++{
		split_input = strings.Split(input[i]," ")
		city[i] = split_input[0]
		diff_time[i],_ = strconv.Atoi(split_input[1])
		
		if s_time_h + diff_time[i] >= 0 && s_time_h + diff_time[i] < 24{
			tmp_h = s_time_h + diff_time[i]
		}else if s_time_h + diff_time[i] < 0{
			tmp_h = s_time_h + diff_time[i] + 24
		}else if s_time_h + diff_time[i] >= 24{
			tmp_h = s_time_h + diff_time[i] - 24
		}
		if tmp_h > 9{
			fmt.Printf("%d:%s\n",tmp_h,base_m)
		}else {
			fmt.Printf("0%d:%s\n",tmp_h,base_m)
		}
	}
	
}
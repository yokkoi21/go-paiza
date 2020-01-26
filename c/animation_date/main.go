package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input string
	var split_input []string
	var split_input_date []string
	var split_input_time []string
	var month,day,hour,minitue int
	var month_s,day_s,hour_s,minitue_s string
	

	sc.Scan()
	input = sc.Text()
	split_input = strings.Split(input," ")
	split_input_date = strings.Split(split_input[0],"/")
	split_input_time = strings.Split(split_input[1],":")
	month,_ = strconv.Atoi(split_input_date[0])
	day,_ = strconv.Atoi(split_input_date[1])
	hour,_ = strconv.Atoi(split_input_time[0])
	minitue,_ = strconv.Atoi(split_input_time[1])

	if hour > 23 {
		day += hour / 24
		hour = hour % 24
	}

	if month < 10{
		month_s = fmt.Sprintf("0%d",month)
	}else {
		month_s = fmt.Sprintf("%d",month)
	}
	if day < 10{
		day_s = fmt.Sprintf("0%d",day)
	}else {
		day_s = fmt.Sprintf("%d",day)
	}
	if hour < 10{
		hour_s = fmt.Sprintf("0%d",hour)
	}else {
		hour_s = fmt.Sprintf("%d",hour)
	}
	if minitue < 10{
		minitue_s = fmt.Sprintf("0%d",minitue)
	}else {
		minitue_s = fmt.Sprintf("%d",minitue)
	}
	
	fmt.Printf("%s/%s %s:%s",month_s,day_s,hour_s,minitue_s)
	
}
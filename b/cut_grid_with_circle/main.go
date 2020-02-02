package main
import "fmt"
import "strings"
import "os"
import "strconv"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [183]string
	var required_time []string
	var train_time [181]string
	var train_time_split []string
	var num int
	var i,j int
	var a,b,c,t_time_h,t_time_m,ns_to_c int
	var c_time_h,c_time_m int
	var latest_work_time int
	var t_time_tmp int

	i = 0
	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	
	required_time = strings.Split(input[0]," ")
	num,_ = strconv.Atoi(input[1])

	j = 2
	for i = 0;i < num;i++{
		train_time[i] = input[j]
		j++
	}
	// fmt.Println(train_time[0])
	// fmt.Println(train_time[1])

	a,_ = strconv.Atoi(required_time[0])
	b,_ = strconv.Atoi(required_time[1])
	c,_ = strconv.Atoi(required_time[2])
	 
	ns_to_c = b + c
	latest_work_time = (8 * 60) + 59
	for i = num;i > -1;i--{
		if train_time[i-1] != ""{
			i_tmp := i - 1
			train_time_split = strings.Split(train_time[i_tmp]," ") 
			t_time_h,_ = strconv.Atoi(train_time_split[0])
			t_time_m,_ = strconv.Atoi(train_time_split[1])
			t_time_tmp = (t_time_h * 60 ) + t_time_m + ns_to_c
			if t_time_tmp <= latest_work_time {
				t_time_tmp = t_time_tmp - a
				t_time_tmp = t_time_tmp - ns_to_c
				break
			}
		}
	}
	
	c_time_h = t_time_tmp / 60
	c_time_m = t_time_tmp % 60
	if c_time_m <= 10{
		fmt.Printf("0%d:0%d",c_time_h,c_time_m)
	}else {
		fmt.Printf("0%d:%d",c_time_h,c_time_m)
	}
	

}
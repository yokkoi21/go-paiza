package main
import "strings"
import "os"
import "strconv"
import "fmt"
import "bufio"

func main(){
	var sc = bufio.NewScanner(os.Stdin)
	var input [1002]string
	var split_input []string
	var base_atk,base_def,base_agi,n,i int
	var monster_name string
	var min_atk_tmp,max_atk_tmp,min_def_tmp,max_def_tmp,min_agi_tmp,max_agi_tmp int
	var ans_check int
	

	for sc.Scan(){
		input[i] = sc.Text()
		i++
	}
	split_input = strings.Split(input[0]," ")
	base_atk,_ = strconv.Atoi(split_input[0])
	base_def,_ = strconv.Atoi(split_input[1])
	base_agi,_ = strconv.Atoi(split_input[2])
	n,_ = strconv.Atoi(input[1])
	
	for i = 2;i < n+2;i++{
		split_input = strings.Split(input[i]," ")
		monster_name = split_input[0]
		min_atk_tmp,_ = strconv.Atoi(split_input[1])
		max_atk_tmp,_ = strconv.Atoi(split_input[2])
		min_def_tmp,_ = strconv.Atoi(split_input[3])
		max_def_tmp,_ = strconv.Atoi(split_input[4])
		min_agi_tmp,_ = strconv.Atoi(split_input[5])
		max_agi_tmp,_ = strconv.Atoi(split_input[6])
		if (min_atk_tmp <= base_atk && max_atk_tmp >= base_atk) {
		    if (min_def_tmp <= base_def && max_def_tmp >= base_def) {
		        if (min_agi_tmp <= base_agi && max_agi_tmp >= base_agi){
			        fmt.Println(monster_name)
			        ans_check += 1
		        }
		    }
		}
	}

	if ans_check < 1{
		fmt.Println("no evolution")
	}
	
	
	
}
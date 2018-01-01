
package main
import "fmt"
import s "strings"


func F3(mask string, line string, k int) (int, bool){
	mask_len:=len(mask); line_len:=len(line); error:=false;
	for j:= 0; j<mask_len; j++{  
		if (mask[j] == '*'){ 
			help := 0; 
			mask_part := mask[j:]
			mask_part_len := len(mask_part)
			full_mask := true; symb := 0;
			for i:=0; i<mask_part_len; i++{
				if mask_part[i] != '*'{
					full_mask = false
					symb = i
					break
				}
			}
			if full_mask{
				k = line_len
				break
			}
			help = s.IndexByte(line[k:], (mask[j+symb]))
			if help != -1{
				k = k + help
				k++
				j= j+symb
			}else{break}
		}else if line[k] != mask[j]{
			error = true; break;
		}else {k++}
	}
	return k, error
}


func Task(line string, mask string) string{
	line_len := len(line); 
	result:=""; 
	for i:=0; i<line_len; i++{  
		start_index:=i; 
		var k, error = F3(mask, line, i)
		if !error{  
			result =result + "  " + line[start_index: k]
		}
	}
	return result
}


func main(){
	fmt.Println(Task("ba bbbac", "ba**")) 
 	  
}



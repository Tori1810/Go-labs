
//package code
package main
import "fmt"

func F1(line string, symb byte) (int, bool){  
	line_len := len(line)
	for i:=0; i<line_len; i++{
		if line[i] == symb{
			return i, false
		}
	} 
	return 0, true
}


func F2(mask_part string) (int, bool){
	mask_part_len := len(mask_part)
	var i int
	for i:=0; i<mask_part_len; i++{
		if mask_part[i] != '*'{
			return i, false
		}
	}
	return i, true
}


func F3(mask string, line string, k int) (int, bool){
	mask_len:=len(mask); line_len:=len(line); full_mask:=false; error:=false;
	for j:= 0; j<mask_len; j++{  
		if (mask[j] == '*'){ 
			help := 0; symb := 0
			symb, full_mask = F2(mask[j:])
			if full_mask == true{
				k = line_len
				break
			}
			help, error = F1(line[k:], mask[j+symb])
			if error == false{
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
		k:=i; 
		error:=false; 
		start_index:=i; 
		k, error = F3(mask, line, k)
		if error == false{  
			result =result + "  " + line[start_index: k]
		}
	}
	return result
}




func main(){
 	fmt.Println(Task("ba bbbac", "b*a**"))   
}



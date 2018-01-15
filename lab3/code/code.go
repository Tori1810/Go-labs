
package code

import "sync"
import s "strings"

var result string = "";


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
	var wg sync.WaitGroup
	line_len := len(line); 

	for i:=0; i<line_len-line_len%5; i+=5{ 
		wg.Add(5)
		go gorutine(i, mask, line,  &wg);
		go gorutine(i+1, mask, line,  &wg);
		go gorutine(i+2, mask, line,  &wg);
		go gorutine(i+3, mask, line,  &wg);
		go gorutine(i+4, mask, line,  &wg);
		wg.Wait()
	}
	for i:=0; i<line_len%5; i++{
		start_index:=line_len - i -1; 
		var k, error = F3(mask, line, line_len - i -1);
		if !error{  
			result =result + "  " + line[start_index: k];
		}
	}
	
	return result
}

func gorutine(i int, mask string, line string,  wg *sync.WaitGroup){
	start_index:=i; 
	var k, error = F3(mask, line, i);
	if !error{  
		result =result + "  " + line[start_index: k];
	}
	wg.Done()
}






//func main(){
 	//fmt.Println(task("ba bbbac", "b*a**"))   
//}



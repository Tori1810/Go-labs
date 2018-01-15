package main

import(	"testing"
 		"io/ioutil"
		"labs/lab2/code"
)

 func TestFull(t *testing.T) {
	
     contents,_ := ioutil.ReadFile("tests.txt")
 	var str = string(contents)
        
     ioutil.WriteFile("output.txt", []byte(code.Task(str, "a*****b*****c")), 0644)
 }

/*func TestShort(t *testing.T) {
	
    contents,_ := ioutil.ReadFile("tests__short.txt")
	var str = string(contents)
        
    ioutil.WriteFile("output__short.txt", []byte(code.Task(str, "a*****b*****c")), 0644)
}*/


/*func TestTask(t *testing.T){
	res := code.Task("adsb fg afihfiuhfpb", "a***b")
	if res != "  adsb  afihfiuhfpb"{
		t.Error( 
            "expected", "  adsb  afihfiuhfpb",
            "got", res,
        )
	}
}*/


/*func TestF1(t *testing.T) {
	res, flag := code.F1("abc", 'c')
	if 2 != res{
        t.Error( 
            "expected", 2,
            "got", res,
        )
    }
    if flag != false{
    	t.Error( 
            "expected", false,
            "got", flag,
        )
    }
}*/


/*func TestF2(t *testing.T){
	res, flag:= code.F2("***b")
	if 3 != res {
        t.Error( 
            "expected", 3,
            "got", res,
        )
    }
    if flag != false{
    	t.Error( 
            "expected", false,
            "got", flag,
        )
    }
}*/

/*func TestF3(t *testing.T){
	res, flag := code.F3("a*d", "abcd", 0)
	if 4 != res {
        t.Error( 
            "expected", 4,
            "got", res,
        )
    }
    if flag != false{
    	t.Error( 
            "expected", false,
            "got", flag,
        )
    }
}*/

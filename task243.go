package main
import ("flag"
		"fmt"
		"strconv"	
		)		
func main(){

	p_m := flag.String("f", "", "Take argument from file")
	flag.Parse()
	var m, err = strconv.Atoi(*p_m)
	if err != nil {
		panic(err)
	}
for x:=1; (x*x)<m; x++ {
	for y:=1; (y*y)<m; y++ {
		if ((x*x)+(y*y)) == m {
			fmt.Println(x, y)
		}
		
	}
	
}

}
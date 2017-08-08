package main

import	(
	"fmt"
	"strconv"
	)

func main()  {

	total:=0
    for i:=0;i <= 15;i++ {
        total+=i
	}
	
	fmt.Println("Total : " + strconv.Itoa(total))
}



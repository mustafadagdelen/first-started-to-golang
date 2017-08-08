package main

import "fmt"

func main() {
	var	result int 
	result=  Topla(15, 80)
	fmt.Println(result)

	var myArray []int
	myArray[0]=4
	myArray[1]=28
	myArray[2]=79
	myArray[3]=12
	myArray[4]=15

	var arraySum int = ArrayTopla(myArray)
	fmt.Println("Toplamı :", arraySum)
}

func Topla(a, b int) int {
	return (a + b);
}

func ArrayTopla(param []int)int {
	var countOfArray int = len(param)
	if(countOfArray <1){
		return 0;
	}
	
	var count int=0
	for i := 0; i < countOfArray; i++ {
		count = count + param[i]
	}

	return count;
}
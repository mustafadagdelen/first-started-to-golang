import ("errors")

func ErrorSample(x int)(int,error)  {
	if(x==0){
		return x,errors.New("Selam ben bir hata")
	}else{
		x++
		return x,nil
	}
}
package main
import (
	"fmt"
	"errors"
)

func main(){
	fmt.Println("Hello World!")
	var intNum int16=7
	fmt.Println(intNum)
	var num int = 10
	var den int =0
	var result, remainder, err = divide(num,den)
	if err!=nil{
		fmt.Printf(err.Error())
	}else{
		fmt.Printf("The result is %v and the remainder is  %v",result,remainder)
	}

}

func divide(numerator int, denominator int) (int,int,error){
	var err error
	if denominator==0{
		err=errors.New("Cannot divide by 0")
		return 0,0,err
	}
	var result int=numerator/denominator
	var remainder int=numerator%denominator
	return result, remainder, err
}
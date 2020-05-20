package main

import (
"fmt"
"math/cmplx"
)


func main(){
	
	cateto1, cateto2 :=1+0i, 0+1i
	fmt.Println("La hipotenusa tiene una longitud igual a %v\n", cmplx.Sqrt(cmplx.Pow(cateto1, 2) + cmplx.Pow(cateto2, 2)))


}

package main

import (
	"fmt"
)

func main() {
	// var firstnum int = 290
	// fmt.Println("The First Number Is : ", firstnum)
	
	// // var secondnum float64 = 12.22
	// // var thirdnum float64 = 42.22
	// // fmt.Println("The Number are : ", secondnum, thirdnum)
	
	// // var secondnum, thirdnum = 12.22, 42.22
	// // fmt.Println("The Number are : ", secondnum, thirdnum)
	
	// // secondnum, thirdnum := 12.22, 42.22
	// // fmt.Println("The Number are : ", secondnum, thirdnum) 

	// fmt.Println("Time Now Is",time.Now())

	// variable1 := 9

	// const variable2 int = 90

	// fmt.Println("The First variable Is : ", variable1)
	// fmt.Println("The First variable Is : ", variable2)

	var (
	 nama string = "angga bagus pratama"
	 umur int	= 19
	 suhu float64 = 24
	 tahun int = 2020 - umur
	 fahrenheit float64 =  suhu * 9/5 +32
	 i float64 = fahrenheit
	)
	fmt.Println("Nama : ",nama)
	fmt.Println("Umur : ",umur)
	fmt.Println("Tahun Lahir : ",tahun)
	fmt.Println("Celcius : ",suhu)
	fmt.Println("Fahrenheit : ",fahrenheit)

	if i < 73 {
		fmt.Println("dingin")
	}else{
		fmt.Println("panas")
	}

	




}

package main

import "fmt"
import "strconv"

func main(){
	defer recovery()
	var input string
	fmt.Print("Masukkan Angka : ")
	fmt.Scanln(&input)
	
	var number int;
	var err error
	
	number, err = strconv.Atoi(input)
	
	if err == nil {
		fmt.Println(number, "Ini adalah angka")
	} else {
		fmt.Println(input, "ini Bukan Angka")
		panic(err.Error())
		fmt.Println("Munculkan Saya")
	}	
}

func recovery(){
	if r:=recover();r!=nil {
		fmt.Println("Pemulihan")
	}
}
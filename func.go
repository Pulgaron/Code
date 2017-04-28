package main

import "fmt"

func main(){
	suma:= func (x int, y int )int {
	return x+y
	}
	fmt.Println("Resultado de la suma:",suma (23,13))
}

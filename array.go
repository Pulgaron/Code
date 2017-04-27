package main

import "fmt"

func main(){
	array := [3]float64 {1.1, 3.2, 1.4}
var total float64 = 0
for _, valor := range array{
	total += valor
}

fmt.Println("total:", total)
}

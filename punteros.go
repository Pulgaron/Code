package main

import "fmt"

type Employee struct{

	Name string 
	LastName string
	Address string 
}


func main(){
	a:= Employee{"John","Smith","5 avenue"}
	fmt.Println(a)
	b:= &Employee{}
	b.Name = "Anna"
	b.LastName = "Frank"
	b.Address = "Frankfurt"
	fmt.Println(b)
}

package main

import "fmt"

func main(){
	if val := suma(); val < 0{
	fmt.Println("Valor menor a cero")
	}else{
		fmt.Println("Valor mayor a cero")
	}

}
func suma () int {
	return 10	
}

package main

import "fmt"

func main() {

	for horas:= 0;horas <= 12; horas++{
		fmt.Print("Horas: ",horas)
		fmt.Print("\n")
	}
	for min := 0; min <= 60; min++{
		fmt.Print("minutos: ",min)
		fmt.Print("\n")
	}
}
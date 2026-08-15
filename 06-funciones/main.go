package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hola", nombre)

}
func sumar(a int, b int) int {
	resultado := a + b
	return resultado

}
func restar(a int, b int)int{
	return a-b
}
func multiplicar(a int, b int)int{
	return a*b
}

func main() {
	a := 10
	b := 5
	fmt.Println(sumar(a, b))
	fmt.Println(restar(a, b))
	fmt.Println(multiplicar(a, b))
	saludar("LEONARDO")
}

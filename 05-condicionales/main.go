package main

import "fmt"

func main() {
	/*
		Los condicionales if/else permiten que un programa tome decisiones dependiendo de si una
		condición es verdadera o falsa. Por ejemplo:
	*/
	edad := 20

	if edad >= 18 {
		fmt.Println("SI ES MAYOR")
	} else {
		fmt.Println("SOY MENOR")
	}
	/*
		Que sucede aquí.
		Establecemos una variable "edad" que se inicializa con el valor de "20". Ahora usamos
		el condicional if y hacemos una comparación para que el programa tome una decisión.
		Si edad es mayor o igual a 18 entoces se ejecuta lo que esta dentro del condicional if,
		es decir se imprime "SI SOY MAYOR". Si no, entramos al "else" (que significa; Caso contrario
		si es menor a 18) e imprimimos "SOY MENOR"

	*/

	fmt.Println()
}

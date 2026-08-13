package main

import "fmt"

func main(){
	/*Ahora es necesario que GO tome deciciones; y para ello utilizaremos los comparadores:
			> Mayor que
			< Menor que
			>= Mayor o Igual
			<= Menor o Igual
			== Igual
			!= Diferente

	*/
	edad := 33

	fmt.Println(edad > 18)  // true
	fmt.Println(edad < 18)  // false
	fmt.Println(edad == 33) // true
	fmt.Println(edad != 33) // false

	//Lo importante de estas comparaciones es que se produce como respuesta un boleano (true o false)


}
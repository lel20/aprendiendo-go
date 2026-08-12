package main

import "fmt"

// AHORA VAMO CON EL TIPO DE DATOS QUE MANEJA GO
/*GO es capas de inferir el tipo de datos que se van a utilizar, es decir,
que al escribir lo siguiente:
*/
func main() {
	nombre := "Leonardo"
	edad := 33
	altura := 1.78
	programador := true

	/*Go es capaz de interpretar el código sin necesidad de que nosostros mismos
	  establescamos el tipo de datos. Internamente GO lo entiende de la siguiente manera:
	  nombre       → string
	  edad         → int
	  altura       → float64
	  programador  → bool
	*/

	fmt.Println(nombre)
	fmt.Println(edad)
	fmt.Println(altura)
	fmt.Println(programador)

	/*Implicitamente también podemos establecer el tipo de variables a utilizar. Pero debemos
	  tener en cuenta lo siguiente, el valor := es el encargado de crear la variable y
	  asignarle el tipo de dato
	*/
	var nombre1 string = "sebastian"
	fmt.Println(nombre1)
	//También se pueden declararar una variable de la siguiente manera:
	var edad1 int
	edad1 = 33
	fmt.Println(edad1)
	/* Hay que tener en cuanta que se utilza "var" cuando se quiere 
	declarar una variable implicitamente sin asignarle inicialmente un valor*/
     

	/*
	Ahora, ¿qué pasa cuando a las variables no se les asigan niguna valor?. Lo que sucede es que 
	GO les asigna un valor zero de acuerdo a su tipo de datos. Por ejemplo, en el siguiente código
	sucede los siguiente:
	*/
	var nombre2 string  //nombre2 --> ""
	var edad2 int       //edad2 --> 0
	var activo2 bool    //activo2 --> false
	var altura2 float64 //altura2 --> 0

	fmt.Println(nombre2)
	fmt.Println(edad2)
	fmt.Println(activo2)
	fmt.Println(altura2)

	// CONSTANTES

		//GO también es capaz de establecer valores que no se pueden cambiar, esto se hace con CONST

		const pais = "España"
		const diasSemana = 7

		fmt.Println(pais)
		fmt.Println(diasSemana)
        // Si se cambia el valor se producira un error 
		diasSemana = 8


}

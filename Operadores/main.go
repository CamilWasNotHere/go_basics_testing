package main

import (
	"fmt"
)

func func_ejemplo() bool {

	return true
}

func main() {

	edad := 19
	// Los operadores nos devuelven un True o un False
	fmt.Println(edad > 17)  // True
	fmt.Println(edad < 17)  // False
	fmt.Println(edad >= 17) // True
	fmt.Println(edad <= 17) // False
	fmt.Println(edad == 17) // False
	fmt.Println()
	// Operador OR

	fmt.Println(edad > 17 || edad == 19) // True
	fmt.Println(edad > 20 || edad == 20) // False
	fmt.Println()

	// Operador AND

	fmt.Println(edad >= 17 && edad == 19) // True
	fmt.Println(edad >= 17 && edad == 20) // False
	fmt.Println()

	// Operador Not

	fmt.Println(!true)  // false
	fmt.Println(!false) // True
	fmt.Println()

	fmt.Println(edad < 20)    // True
	fmt.Println(!(edad < 20)) // False
	fmt.Println()
	// condicionales

	if edad >= 19 { // Si se cumple la condición se imprime el mensaje

		fmt.Printf("Acepta\n")
		fmt.Println()
	}

	boolean := false

	if boolean {

		fmt.Println("Es verdadero")
	} else {

		fmt.Println("Es falso")
		fmt.Println()
	}

	if value := func_ejemplo(); value { // Forma de hacer condiconales con funciones

		fmt.Println("Retorno true")
	}

}

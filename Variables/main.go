package main // definimos la función princiapal

import ( // importamos los pquetes que vamos a usar
	"fmt" // fmt es un paquete para formatear y mostrar texto en consola
	"unsafe"
)

func main() {

	var variable_entera int = 12 // Si defenimos el tipo de variable como int el valor puede ser valores de - infinito hasta + infinito
	fmt.Println("El valor de la variable es: ", variable_entera)

	variable_entera2 := 15 // Una forma similar a la sintaxis de python para declarar variables, el tipo de dato lo define el valor de la variable
	fmt.Println("El valor de la variable es: ", variable_entera2)

	var variable_uint uint = 12 // Si definimos el tipo de variable como unit lo valores tienen que ser + infinitivo
	fmt.Println("el valor de la variable es: ", variable_uint)

	var variable_string string = "Illojuan" // tipo de dato string
	fmt.Println("El nombre de la persona es: ", variable_string)

	variable_string2 := "Golang :)" // otra forma de definir un tipo de dato string
	fmt.Println("El nombre de la persona es: ", variable_string2)

	var variable_boleana bool = true                                // tipo de dato boleano
	fmt.Printf("El valor de el boleano es: %t\n", variable_boleana) //formatenado los datos con Printf

	variable_boleana2 := false // otra forma de definir el tipo de dato boleano
	fmt.Printf("El valor de el boleano es: %t\n", variable_boleana2)

	fmt.Println("Ver la direccion de memoria de una variable", &variable_boleana2) // vemos la dirección de memoria utlizando ampersand & y la variable

	const constante int = 12 // Definiendo una constante con un valor entero
	fmt.Println("La constante vale: ", constante)

	const constante2 = 12 // Otra forma de definir una constante
	fmt.Println("La constante vale: ", constante2)

	const constante_string string = "OAAAA"
	println("El valor de la constante es: ", constante_string) // constante de tipo string

	//Acontinuación tipos de datos enteros en diferentes bits
	var my_byte8 int8
	fmt.Printf("Type: %T byte: %d bytes: %d\n", my_byte8, unsafe.Sizeof(my_byte8), unsafe.Sizeof(my_byte8)*8)

	var my_byte16 int16
	fmt.Printf("Type: %T byte: %d bytes: %d\n", my_byte16, unsafe.Sizeof(my_byte16), unsafe.Sizeof(my_byte16)*8)

	var my_byte32 int32
	fmt.Printf("Type: %T byte: %d bytes: %d\n", my_byte32, unsafe.Sizeof(my_byte32), unsafe.Sizeof(my_byte32)*8)

	var my_byte64 int64
	fmt.Printf("Type: %T byte: %d bytes: %d\n", my_byte64, unsafe.Sizeof(my_byte64), unsafe.Sizeof(my_byte64)*8)

	var my_uint64 uint64 // Ejemplo del uso de uint con 64 bits aplicable para los otros bits
	fmt.Printf("Tpye: %T byte: %d bytes: %d\n", my_uint64, unsafe.Sizeof(my_uint64), unsafe.Sizeof(my_uint64)*8)

	var my_float32 float32 // Ejemplo dato de tipo flotante de 32 bits
	fmt.Printf("Tpye: %T byte: %d bytes: %d\n", my_float32, unsafe.Sizeof(my_float32), unsafe.Sizeof(my_float32)*8)

	var my_float64 float64 // Ejemplo dato de tipo flotante 64 bits
	fmt.Printf("Type %T byte: %d bytes: %d\n", my_float64, unsafe.Sizeof(my_float64), unsafe.Sizeof(my_float64)*8)

	var my_float float64 // Ejemplo del valor predefinido de un tipo flotante si no se le agrega ningun valor
	fmt.Printf("Valor predefinido del tipo flotante %f\n", my_float)

	my_string := `Esta es una forma de escribir texto
con saltos de linea sin 
ningun problema en el codigo` //utilizamos `` para escribir texto con saltos de linea :)
	fmt.Printf("%s", my_string)

}

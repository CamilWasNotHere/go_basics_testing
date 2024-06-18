package main // definimos la función princiapal

import ( // importamos los pquetes que vamos a usar

	"fmt"     // fmt es un paquete para formatear y mostrar texto en consola
	"strconv" // convertir string a un tipo de dato
	"unsafe"  // unsafe es un paquete que se utiliza para saber el tamaño de memoria que una variable ocupa en byte
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

	{ //definimos un Scope de tipo local
		fmt.Println() // Salto de linea
		my_float := 32.3123
		fmt.Printf("Type: %T valor: %f\n", my_float, my_float) // Vemos el tipo de dato del cual estamos tratando

		my_float_string := fmt.Sprintf("%f", my_float) // formatemos el tipo de dato florante y lo convertimos a un string
		fmt.Printf("Type: %T valor: %s\n", my_float_string, my_float_string)

		my_int := 12
		fmt.Printf("Type %T, valor: %d\n", my_int, my_int) // vemos el tipo de dato del cual estamos tratando

		my_int_string := fmt.Sprintf("%d", my_int) // formatiamos el tipo de dato entero y lo convertimos a un string
		fmt.Printf("Type: %T, valor: %s\n", my_int_string, my_int_string)

		my_string_to_int, err := strconv.ParseInt("42", 0, 64) // Forma de poder convertir un string a un entero
		fmt.Printf("Type: %T, valor: %d\n", my_string_to_int, my_string_to_int)
		fmt.Println(err)

		my_string_to_float, _ := strconv.ParseFloat("12.23", 64)                      // Forma de convertir un string a un flotante, utilizamos _ para variables que no vamos a utilizar
		fmt.Printf("Type: %T, valor: %.2f\n", my_string_to_float, my_string_to_float) //para redondear utilizar un . y la cantidad de decimales seguido de f

	}

	{

		var A byte = 65
		fmt.Println("Valor ASCII de A: ", A) //convirtiendo el valor en byte a un string en este caso a 'A'

		fmt.Println(string(rune(('👻')))) // El tipo de codificacion rune o unicode acepta ermogis

		var vector = []byte{65, 96, 88, 115} // Definimos un vector de tipo byte con valores que convertiremos en caracteres
		fmt.Println(vector)                  // Imprimos el vector con sus valores
		fmt.Println(string(vector))          // Convertimos y imprimos el vector a caracteres

	}

	{

	}
}

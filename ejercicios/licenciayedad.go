////

package main

import (
	"fmt"
)

func main() {

	edad := 19
	licencia := true

	if edad > 15 && licencia {

		fmt.Println("Puede seguir avanzando")

	} else {

		fmt.Println("No puede seguir circulando")
	}

}

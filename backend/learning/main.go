package main

import (
	"fmt"
	"strings"
)

//var entero int = 12

func main() {
	entero := 13
	decimal := 12.5
	suma := float64(entero) + (decimal)

	fmt.Println(suma)
	mensaje := "Hola, "
	nombre := "Omar"

	concatenado := mensaje + nombre
	concatenado = strings.ToUpper(concatenado)
	fmt.Println(concatenado)

	arrayFijo := [2]int{1, 2}
	arrayDinamico := []int{1, 2, 3, 4, 5}
	fmt.Println(arrayFijo)
	fmt.Println(arrayDinamico)

	arrayDinamico = append(arrayDinamico, 6)

	//maps
	diccionario := map[string]string{
		"nombre": "Omar",
		"edad":   "20",
		"ciudad": "Colombia",
	}
	fmt.Println(diccionario)

	type Persona struct {
		Nombre string
		Edad   int
		Ciudad string
	}

	persona := Persona{
		Nombre: diccionario["nombre"],
		Ciudad: "Colombia",
	}

	fmt.Println(persona)
	fmt.Println(persona.Nombre)

}

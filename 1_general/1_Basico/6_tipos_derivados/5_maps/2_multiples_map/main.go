/*
Ejemplo Como declarar variables
En este ejemplo intento mostrar
como usar map instanciarlos imprimirlos recorrerlos
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("-------------------1---------------------------")
	var map_1 map[int]int

	// Checking if the map is nil or not
	if map_1 == nil {

		fmt.Println("True")
	} else {

		fmt.Println("False")
	}
	fmt.Println("-------------------2---------------------------")
	// Creating and initializing a map
	// Using shorthand declaration and
	// using map literals
	map_2 := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	fmt.Println("Map-2: ", map_2)
	fmt.Println("-------------------3---------------------------")
	// Creating a map
	// Using make() function
	var My_map = make(map[float64]string)
	fmt.Println(My_map)

	// As we already know that make() function
	// always returns a map which is initialized
	// So, we can add values in it
	My_map[1.3] = "Rohit"
	My_map[1.5] = "Sumit"
	fmt.Println(My_map)
	fmt.Println("-------------------4---------------------------")
	// Creating and initializing a map
	m_a_p := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	// Iterating map using for rang loop
	for id, pet := range m_a_p {

		fmt.Println(id, pet)
	}
	fmt.Println("-------------------5---------------------------")
	// Creating and initializing a map
	m_a_p = map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	fmt.Println("Original map: ", m_a_p)

	// Adding new key-value pairs in the map
	m_a_p[95] = "Parrot"
	m_a_p[96] = "Crow"
	fmt.Println("Map after adding new key-value pair:\n", m_a_p)

	// Updating values of the map
	m_a_p[91] = "PIG"
	m_a_p[93] = "DONKEY"
	fmt.Println("\nMap after updating values of the map:\n", m_a_p)
}

package main

/*
Ejemplo slice
En este ejemplo intento mostrar
como usar la sintaxis de low:high:max
en los slice
*/
import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	//a[low : high]
	var ejemplo1 []int = primes[0:5]

	var ejemplo2 []int = primes[:3]
	var ejemplo3 []int = primes[3:]
	//a[low : high: max]
	var ejemplo4 []int = primes[1:3:6]
	var ejemplo5 []int = primes[0:3:5]
	var ejemplo6 []int = primes[0:2:2]
	fmt.Println("-----------ejemplo1--------------")
	fmt.Println(ejemplo1)
	fmt.Println("len:", len(ejemplo1))
	fmt.Println("capacity:", cap(ejemplo1))
	fmt.Println("-----------ejemplo2---------------")
	fmt.Println(ejemplo2)

	fmt.Println("len:", len(ejemplo2))
	fmt.Println("capacity:", cap(ejemplo2))
	fmt.Println("-----------ejemplo3---------------")
	fmt.Println(ejemplo3)

	fmt.Println("len:", len(ejemplo3))
	fmt.Println("capacity:", cap(ejemplo3))
	fmt.Println("-----------ejemplo4---------------")
	fmt.Println(ejemplo4)

	fmt.Println("len:", len(ejemplo4))
	fmt.Println("capacity:", cap(ejemplo4))
	fmt.Println("-----------ejemplo5---------------")
	fmt.Println(ejemplo5)

	fmt.Println("len:", len(ejemplo5))
	fmt.Println("capacity:", cap(ejemplo5))
	fmt.Println("-----------ejemplo6---------------")
	fmt.Println(ejemplo6)

	fmt.Println("len:", len(ejemplo6))
	fmt.Println("capacity:", cap(ejemplo6))
}

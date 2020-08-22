/*
Ejemplo de operadores logico
En este ejemplo
se muestra los operadores logicos y operaciones de corridas de bit
*/

package main

import "fmt"

func main() {
	x := true  //bool
	y := false //bool
	fmt.Println(x || y)
	fmt.Println(x && y)
	fmt.Println(x && !y)
	fmt.Println(0 ^ 1)
	fmt.Println("-----------Corridas de bits ")
	fmt.Println(001 << 1) // 010
	fmt.Println(001 << 2) //100
	fmt.Println(001 << 3) //1000
	fmt.Println(001 << 4) //10000

	fmt.Println("-----------corridas de bits ¿por que esta mal?")
	fmt.Println(10 >> 1)    // 010
	fmt.Println(100 >> 2)   //100
	fmt.Println(1000 >> 3)  //1000
	fmt.Println(10000 >> 4) //10000

	fmt.Println("-----------corridas de bits ")
	fmt.Println(2 >> 1)  // 010
	fmt.Println(4 >> 2)  //100
	fmt.Println(8 >> 3)  //1000
	fmt.Println(16 >> 4) //10000

	fmt.Println("-----------que numero necesitamos? ")
	fmt.Println(1 >> 1) // 2
	fmt.Println(1 >> 2) //4
	fmt.Println(1 >> 3) //8
	fmt.Println(1 >> 4) //16

	////////////////////////////////////////
	/*  16|8|4|2|1
	    0|0|0|0|0 =  0+0+0+0+0 = 0
	    0|0|0|0|1 =  0+0+0+0+1 = 1
	    0|0|0|1|0 =  0+0+0+2+0 = 2
	    0|0|0|1|1 =  0+0+0+2+1 = 3
	    0|0|1|0|0 =  0+0+4+0+0 = 4
	    0|0|1|0|1 =  0+0+4+0+1 = 5
	    0|0|1|1|0 =  0+0+4+1+0 = 6
		0|0|1|1|1 =  0+0+4+2+1 = 7
		0|1|0|1|0 =  0+8+0+2+0 = 10


	    1|1|0|0|0 =  16+8+0+0+0 = 24
	         .
	         .
	         .
	    1|0|0|0|0 =  16+0+0+0+0 = 16
	         .
	         .
	         .
	    1|1|1|1|1 =  16+8+4+2+1 =31

	*/
}

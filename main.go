package main

import (
	"fmt"
)

/*
	func main() {
		a := rand.Intn(10)
		b := rand.Intn(10)

		fmt.Println(a)
		fmt.Println(b)

			if (a < 4) && (b < 4) {
				fmt.Println("The numbers for 4 are", a, b)
			} else if (a > 6) && (b > 6) {
				fmt.Println("The numbers for 6 are", a, b)
			} else if (4 <= a) && (a <= 6) {
				fmt.Println("The numbers for 4 and 6 are", a)
			} else if b != 5 {
				fmt.Println("The number for not 5 are", b)
			} else {
				fmt.Println("none of the previous were met")
			}
		}

		switch {
		case (a < 4) && (b < 4):
			fmt.Println("The numbers for 4 are", a, b)
		case (a > 6) && (b > 6):
			fmt.Println("The numbers for 6 are", a, b)
		case (4 <= a) && (a <= 6):
			fmt.Println("The numbers for 4 and 6 are", a)
		case b != 5:
			fmt.Println("The number for not 5 are", b)
		default:
			fmt.Println("none of the previous were met")
		}
	}

//another execrise

	func main() {
		//first test
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}

		//second test

		for j := 0; j < 101; j++ {

			a := rand.Intn(10)
			b := rand.Intn(10)

			switch {
			case (a < 4) && (b < 4):
				fmt.Println("The numbers for 4 are", a, b)
			case (a > 6) && (b > 6):
				fmt.Println("The numbers for 6 are", a, b)
			case (4 <= a) && (a <= 6):
				fmt.Println("The numbers for 4 and 6 are", a)
			case b != 5:
				fmt.Println("The number for not 5 are", b)
			default:
				fmt.Println("none of the previous were met")
			}
		}

}

// Another exersice

	func main() {
		for a := 0; a < 43; a++ {
			fmt.Println("iteration", a)

			x := rand.Intn(5)
			switch {
			case x == 0:
				fmt.Printf("%v and %T\n", x, x)
			case x == 1:
				fmt.Printf("%v and %T\n", x, x)
			case x == 2:
				fmt.Printf("%v and %T\n", x, x)
			case x == 3:
				fmt.Printf("%v and %T\n", x, x)
			default:
				fmt.Printf("%v and %T\n", x, x)
			}
		}
	}
	//condition loop
	func main() {
		i := 20
		for i >= 5 {
			fmt.Println(i)
			i--
		}

	}


	// break statement
	func main() {
		i := 0
		for {
			//i := 0
			i++
			fmt.Println(i)
			if i >= 10 {
				break
			}
		}
	}

	// continue statement
	func main() {
		for i := 0; i < 20; i++ {
			if (i % 2) != 0 {
				fmt.Printf("%v is an Odd number\n", i)
				continue
			}
		}
	}

	//nested loop
	func main() {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				fmt.Printf("The first loop %v \t The second loop, number %v\n", i, j)
			}
		}
	}

	//loop throught range with slice data structure
	func main() {
		xi := []int{42, 43, 44, 45, 46, 47}
		for i, v := range xi {
			fmt.Printf("index %v \t value %v\n", i, v)
		}
	}

	//loop through key and value in a map data structure
	func main() {
		m := map[string]int{
			"James":      42,
			"Moneypenny": 32,
		}
		for key, value := range m {
			fmt.Printf("Key is %v, value is %v\n", key, value)
		}

	}

	//new execrise

	func main() {
		m := map[string]int{
			"James":      42,
			"Moneypenny": 32,
		}

		for k, v := range m {
			fmt.Printf("%v and %T", k, v)

			age1 := m["James"]
			fmt.Println("The age of Bond", age1)

			//this line assinged an empty value zero to the var Q
			age2 := m["Q"]
			fmt.Println("The age of Q", age2)

			// this is the idiom ok, we use it to check for values in the map. note:- the "!ok" means not ok, normaly it is "ok"
			if v, ok := m["Q"]; !ok {
				fmt.Println("There is no Q, and here is the zero value of int", v)
			}

		}

	}

	// statement statement idiom
	func main() {
		c := 1
		for i := 0; i < 100; i++ {
			if x := rand.Intn(5); x == 3 {
				fmt.Printf("Iteration %v\t Total count is %v\t x is %v\n", i, c, x)
				c++
			}
		}
	}
*/

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

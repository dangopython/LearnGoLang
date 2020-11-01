package main

import "fmt"

func numbers() {
	var a int = 4
	var b, c int
	b = 5
	c = 10
	fmt.Println(a)
	fmt.Println(b + c)

	var d float64 = 9.14
	fmt.Println(d)
}

func strings() {
	var s string = "This is string s"
	fmt.Println(s)
}

func shorthand() {
	a := 9        // var a int = 9
	b := "golang" // var b string = "golang"
	c := 4.17
	d := false
	e := "Hello"
	f := `Do you like golang, so far?`
	g := 'M'

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}

func exercise() {
	/*
		print following to console
		Sagar Pohekar
		24
		154.61
	*/
	var name string = "Sagar Pohekar"
	var number int = 24
	var length float64 = 154.61
	fmt.Println(name)
	fmt.Println(number)
	fmt.Println(length)
}

func main() {
	fmt.Println("numbers...")
	numbers()
	fmt.Println("strings...")
	strings()
	fmt.Println("shorthand...")
	shorthand()
	fmt.Println("exercise...")
	exercise()

}

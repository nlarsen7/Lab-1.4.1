//Nicholas Larsen
//2/6/2020
//practice on user inputs
package main

import( "fmt" 
)

func main() {
var favfood = "Pizza"
var name string
var age float64

fmt.Println("What is your name?")

fmt.Scanln(&name)

fmt.Println("What is your age?")

fmt.Scanln(&age)

fmt.Println("Welcome", name)

fmt.Println("At your age of", age, "I loved eating", favfood, "too!")
}
package main

import "fmt"

func main() {
	exercise()
}

func exercise() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	bob := Employee{"Bob", "Smith", 127}
	fmt.Println(bob)

	will := Employee{firstName: "Will", lastName: "Riser", id: 4876}
	fmt.Println(will)

	var jim Employee
	jim.firstName = "Jim"
	jim.lastName = "Doe"
	jim.id = 1234
	fmt.Println(jim)

}

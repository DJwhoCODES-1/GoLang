package main

import "fmt"

type Contact struct {
	Email  string
	Mobile string
}

type Address struct {
	House    string
	Block    string
	Landmark string
}

type Person struct {
	firstName string
	lastName  string
	age       int
	isAlive   bool
}

type Employee struct {
	Person_Details  Person
	Person_Contacts Contact
	Person_Address  Address
}

func main() {
	fmt.Println("Struct")

	var Devanshu Person
	Devanshu.firstName = "Devanshu"
	Devanshu.lastName = "Jain"
	Devanshu.age = 24
	Devanshu.isAlive = true
	fmt.Println(Devanshu)

	person1 := Person{
		firstName: "Doraemon",
		lastName:  "Choudhary",
		age:       200,
		isAlive:   true,
	}

	fmt.Println(person1)

	person2 := new(Person)
	person2.firstName = "Shinchan"
	person2.lastName = "Jaiswal"
	person2.age = 5
	person2.isAlive = true

	fmt.Println(person2)

	// emp1 := new(Employee)
	var emp1 Employee
	emp1.Person_Details.firstName = "Raju"
	emp1.Person_Details.lastName = "Rastogi"
	emp1.Person_Details.age = 24
	emp1.Person_Details.isAlive = true

	emp1.Person_Contacts = Contact{
		Email:  "johndoe@gmail.com",
		Mobile: "1234567898",
	}

	emp1.Person_Address = Address{
		House:    "string",
		Block:    "string",
		Landmark: "string",
	}

	fmt.Println(emp1)
}

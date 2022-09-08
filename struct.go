package main

import "fmt"

type Person struct {
	name, address string
	id, age       int
}

type Customer struct {
	name, address string
	id, age       int
}

func test(person Person) {
	person.address = "dsddd"
	person.name = "ki"
	fmt.Printf("person: %v\n", person)
}

func test2(person *Person) {
	person.address = "dsddd"
	person.name = "ki"
	fmt.Printf("person: %v\n", person)
}

func main() {
	var student Person
	student.id = 1
	student.name = "不放弃"
	student.address = "home"
	student.age = 36
	fmt.Printf("student: %v\n", student)

	var cus Customer
	fmt.Printf("student: %v\n", cus)

	var tom Person
	tom = Person{
		id:      1,
		name:    "chunlai",
		address: "home",
		age:     35,
	}

	fmt.Printf("tom: %v\n", tom)

	var tom1 Person
	tom1 = Person{
		"chunlai",
		"home",
		100,
		35,
	}
	fmt.Printf("tom: %v\n", tom1)

	var p_person *Person
	p_person = &tom

	fmt.Printf("p_person: %p\n", p_person)
	fmt.Printf("p_person: %v\n", *p_person)

	//匿名结构体
	var animal struct {
		id   int
		name string
	}

	animal.id = 100
	animal.name = "大熊猫"
	fmt.Printf("animal: %v\n", animal)

	test(tom)

	// var tom2 *Person
	// tom2 = new(Person)
	tom2 := Person{
		address: "bbb",
		id:      10000,
	}
	fmt.Println(tom2)
	// tom2.address = "aaa"
	// tom2.id = 1000
	test2(&tom2)

	fmt.Println(tom2)
}

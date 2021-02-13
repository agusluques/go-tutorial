package main

import "fmt"

type contactInfo struct {
	email string
	phone int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p *person) howToContact() string {
	return "You can reach " + p.firstName + " by calling to " + fmt.Sprint(p.contact.phone) + " or email to " + p.contact.email
}

func (p *person) updateLastName(name string) {
	// p.lastName === (*p).lastName
	p.lastName = name
}

func main() {
	me := person{"Agustin", "Luques", contactInfo{}}

	mom := person{firstName: "Victoria", lastName: "Mingote"}
	mom.firstName = "Maria Victoria"

	var sister person
	sister.firstName = "Luisina"

	fmt.Println(me)
	fmt.Println(mom)
	fmt.Printf("%+v\n", sister)

	me.contact = contactInfo{
		email: "luquesagustin@gmail.com",
		phone: 1133504834,
	}

	fmt.Printf("%+v\n", me)

	fmt.Println(me.howToContact())

	sister.updateLastName("Luques")
	fmt.Println(sister)
}

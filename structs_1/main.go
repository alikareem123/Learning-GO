package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	ali := person{firstName: "Ali", lastName: "Kareem"}
	fmt.Println(ali)

	var Kareem person
	fmt.Println(Kareem)
	fmt.Printf("%+v", Kareem)
	fmt.Println()
	jim := person{
		firstName: "Jim",
		lastName:  "Carrey",
		contact: contactInfo{
			email:   "Jim@carrey.com",
			zipcode: 110023,
		},
	}
	jimPointer := &jim
	jimPointer.updateName2("Sabhash")
	//jim.updateName("Alisson")
	jim.print()
	jimPointer.print()
	mySlice := []string{"bob", "marley"}
	fmt.Println(mySlice)
}

func (p *person) updateName(newFirstName string){
	p.firstName = newFirstName
}
func (p *person) updateName2(newFirstName string){
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

package main

import "fmt"

type Product interface {
	use()
}

type IDCard struct {
	owner string
}

func (c *IDCard) use() {
	fmt.Println("use " + c.owner + "'s ID card")
}

type Factory interface {
	Create(owner string) Product
}

type IDCardFactory struct {
}

func (f *IDCardFactory) Create(owner string) Product {
	return &IDCard{owner: owner}
}

func main() {
	var factory Factory
	factory = &IDCardFactory{}
	card := factory.Create("John Doe")
	card.use()
}

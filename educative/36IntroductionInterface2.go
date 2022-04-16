package main

import (
	"fmt"
)

type User36 struct {
	FirstName, LastName string
}

func (u *User36) Name() string { //Name method used for type User36
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Customer struct {
	Id       int
	FullName string
}

func (c *Customer) Name() string { //Name method used for type Customer
	return c.FullName
}

type INamer interface {
	Name() string //Both Name() methods can be called using the INamer interface
}

func Greet36(n INamer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}
func main() {
	u := &User36{"Matt", "Aimonetti"}
	fmt.Println(Greet36(u))
	c := &Customer{42, "Francesc"}
	fmt.Println(Greet36(c))
}

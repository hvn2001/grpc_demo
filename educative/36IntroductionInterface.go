package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

type Namer interface {
	Name() string //The N_amer interface is defined by the Name() method
} // --> User is Namer interface

func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s", n.Name())
}
func main() {
	u := &User{"Matt", "Aimonetti"}
	fmt.Println(u.Name())
	fmt.Println(Greet(u))
}

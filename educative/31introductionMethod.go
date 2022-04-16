package main

import (
	"fmt"
)

type User31 struct {
	FirstName, LastName string
}

func (u User31) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}
func main() {
	u := User31{"Matt", "Aimonetti"}
	fmt.Println(u.Greeting())
}

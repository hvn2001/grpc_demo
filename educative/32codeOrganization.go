package main

// list of packages to import
import (
	"fmt"
)

// list of constants
const (
	ConstExample = "const before vars"
)

// list of variables
var (
	ExportedVar    = 42
	nonExportedVar = "so say we all"
)

// Main type(s) for the file,
// try to keep the lowest amount of structs per file when possible.
type User32 struct {
	FirstName, LastName string
	Location            *UserLocation
}
type UserLocation struct {
	City    string
	Country string
}

// List of functions
func NewUser(firstName, lastName string) *User32 {
	return &User32{FirstName: firstName,
		LastName: lastName,
		Location: &UserLocation{
			City:    "Santa Monica",
			Country: "USA",
		},
	}
}

// List of methods
func (u *User32) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}
func main() {
	us := User32{FirstName: "Matt",
		LastName: "Damon",
		Location: &UserLocation{
			City:    "Santa Monica",
			Country: "USA"}}
	fmt.Println(us.Greeting())
}

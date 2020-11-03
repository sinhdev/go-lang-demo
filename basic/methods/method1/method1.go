package method1

import "fmt"

//User struct
type User struct {
	FirstName, LastName string
}

//Greeting method
func (u User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

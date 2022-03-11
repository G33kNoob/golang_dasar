package dataGroup

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
	IsAdmin   bool
}

func NewUser() {
	user := User{}
	user.ID = 1
	user.FirstName = "ahsan"
	user.LastName = "abdillah"
	user.Email = "ahsandev2019@gmail.com"
	user.IsActive = true
	user.IsAdmin = true
	fmt.Println(user)
	user2 := User{2, "zahwa", "nurul", "karimah@gmail.com", true, false}
	fmt.Println(user2)
	displayUser1 := displayUser(user2)
	fmt.Println(displayUser1)
}

func displayUser(user User) string {
	result := fmt.Sprintf("list data %s %s email:  %s", user.FirstName, user.LastName, user.Email)
	return result
}

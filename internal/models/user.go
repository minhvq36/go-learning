package models

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) User {
	return User{Name: name, Age: age}
}

func (u User) Greet() {
	fmt.Printf("Hello, %s\n", u.Name)
}

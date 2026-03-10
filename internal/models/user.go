package models

import "fmt"

type User struct {
	Name string
	Age  int
}

func Greet(u User) {
	fmt.Printf("Hello, %s\n", u.Name)
}

func NewUser(name string, age int) User {
	return User{Name: name, Age: age}
}

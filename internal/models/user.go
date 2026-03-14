package models

import "fmt"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func (u User) Greet() {
	fmt.Printf("Hello, %s\n", u.Name)
}

func (u *User) IncreaseAge() {
	u.Age++
}

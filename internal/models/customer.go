package models

import "fmt"

type Customer struct {
	Id       int
	Login    string
	Password string
	Role     string
	Name     string
	Surname  string
	Age      int
}

func (c Customer) String() string {
	return fmt.Sprintf("id: %d, login: %s, password: %s, role: %s, name: %s, surname: %s, age: %d",
		c.Id, c.Login, c.Password, c.Role, c.Name, c.Surname, c.Age)
}

package models

import "fmt"

type Customer struct {
    Id       int    `json:"id"`
    Login    string `json:"login"`
    Password string `json:"password"`
    Role     string `json:"role"`
    Name     string `json:"name"`
    Surname  string `json:"surname"`
    Age      int    `json:"age"`
}
func (c Customer) String() string {
	return fmt.Sprintf("id: %d, login: %s, password: %s, role: %s, name: %s, surname: %s, age: %d",
		c.Id, c.Login, c.Password, c.Role, c.Name, c.Surname, c.Age)
}

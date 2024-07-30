package json

import (
	"encoding/json"
	"fmt"
	"go_psql/internal/models"
	"os"
)

func ReadCustomersFromFile(filename string) ([]models.Customer, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var customers []models.Customer
	err = json.NewDecoder(file).Decode(&customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func WriteCustomersToFile(filename string, customers []models.Customer) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(customers)
}

func GetPersonWithLoginAndPassword(login, password string) (models.Customer, error) {
	path, _ := os.Getwd()
	fmt.Println("path: ", path)
	customers, err := ReadCustomersFromFile("internal/database/json/users.json")
	if err != nil {
		return models.Customer{}, err
	}

	for _, c := range customers {
		if c.Login == login && c.Password == password {
			return c, nil
		}
	}

	return models.Customer{}, fmt.Errorf("customer with login %s and password %s not found", login, password)
}

// GetCustomerByID возвращает клиента по его ID
func GetPerson(id int) (models.Customer, error) {
	path, _ := os.Getwd()
	fmt.Println("path: ", path)
	customers, err := ReadCustomersFromFile("internal/database/json/users.json")
	if err != nil {
		return models.Customer{}, err
	}

	for _, c := range customers {
		if c.Id == id {
			return c, nil
		}
	}

	return models.Customer{}, fmt.Errorf("customer with ID %d not found", id)
}

// CreateCustomer добавляет нового клиента в базу данных
func InsertCustomer(customer models.Customer) error {
	path, _ := os.Getwd()
	fmt.Println("path: ", path)
	customers, err := ReadCustomersFromFile("internal/database/json/users.json")
	if err != nil {
		return err
	}

	customers = append(customers, customer)

	return WriteCustomersToFile("internal/database/json/users.json", customers)
}

// DeleteCustomer удаляет клиента из базы данных по его ID
func RemoveCustomer(id int) error {
	path, _ := os.Getwd()
	fmt.Println("path: ", path)
	customers, err := ReadCustomersFromFile("internal/database/json/users.json")
	if err != nil {
		return err
	}

	var updatedCustomers []models.Customer
	for _, c := range customers {
		if c.Id != id {
			updatedCustomers = append(updatedCustomers, c)
		}
	}

	return WriteCustomersToFile("internal/database/json/users.json", updatedCustomers)
}
func UpdateCustomer(user models.Customer) error {
	path, _ := os.Getwd()
	fmt.Println("path: ", path)
	users, err := ReadCustomersFromFile("internal/database/json/users.json")
	if err != nil {
		return err
	}

	// Ищем пользователя по ID
	for i, u := range users {
		if u.Id == user.Id {
			// Обновляем данные пользователя
			users[i] = user
			return WriteCustomersToFile("internal/database/json/users.json", users)
		}
	}

	return fmt.Errorf("user with ID %d not found", user.Id)
}

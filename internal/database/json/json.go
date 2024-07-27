package json

import (
	"encoding/json"
	"fmt"
	"go_psql/internal/models"
	"io/ioutil"
	"os"
)

func UnmarshalTickets(filepath string) ([]models.Ticket, error) {
	path, err := os.Getwd()
	fmt.Println("path:", path)

	file, err := os.Open("/internal/database/json/tickets.json")
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return []models.Ticket{}, err
	}
	defer file.Close()

	tickets := make([]models.Ticket, 0, 10)
	bs, err := ioutil.ReadAll(file)
	if err != nil {
		return []models.Ticket{}, err
	}

	err = json.Unmarshal(bs, &tickets)
	if err != nil {
		return []models.Ticket{}, err
	}

	return tickets, nil
}

func MarshalTickets(filename string, data []models.Ticket) error {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, bs, 0644)
	if err != nil {
		return err
	}

	return nil
}

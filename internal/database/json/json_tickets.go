package json

import (
	"encoding/json"
	"fmt"
	"go_psql/internal/models"
	"io/ioutil"
	"os"
)

// GetTicketByID возвращает билет по его ID
func GetTicketByID(filepath string, ticketID string) (*models.Ticket, error) {
	tickets, err := UnmarshalTickets(filepath)
	if err != nil {
		return nil, err
	}

	for _, ticket := range tickets {
		if ticket.Id == ticketID {
			return &ticket, nil
		}
	}

	return nil, fmt.Errorf("ticket with ID %s not found", ticketID)
}

// AddTicket добавляет новый билет в файл
func AddTicket(filepath string, ticket models.Ticket) error {
	tickets, err := UnmarshalTickets(filepath)
	if err != nil {
		return err
	}

	tickets = append(tickets, ticket)

	return MarshalTickets(filepath, tickets)
}

// UpdateTicket обновляет информацию о билете
func UpdateTicket(filepath string, ticket models.Ticket) error {
	tickets, err := UnmarshalTickets(filepath)
	if err != nil {
		return err
	}

	for i, t := range tickets {
		if t.Id == ticket.Id {
			tickets[i] = ticket
			return MarshalTickets(filepath, tickets)
		}
	}

	return fmt.Errorf("ticket with ID %s not found", ticket.Id)
}

// DeleteTicket удаляет билет по его ID
func DeleteTicket(filepath string, ticketID string) error {
	tickets, err := UnmarshalTickets(filepath)
	if err != nil {
		return err
	}

	for i, ticket := range tickets {
		if ticket.Id == ticketID {
			tickets = append(tickets[:i], tickets[i+1:]...)
			return MarshalTickets(filepath, tickets)
		}
	}

	return fmt.Errorf("ticket with ID %s not found", ticketID)
}

// UnmarshalTickets читает билеты из файла
func UnmarshalTickets(filepath string) ([]models.Ticket, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tickets := make([]models.Ticket, 0, 10)
	bs, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bs, &tickets)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

// MarshalTickets записывает билеты в файл
func MarshalTickets(filename string, data []models.Ticket) error {
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

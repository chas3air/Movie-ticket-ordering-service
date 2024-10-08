package psql

import (
	"log"

	"go_psql/internal/config"
	"go_psql/internal/models"

	_ "github.com/lib/pq"
)

func GetPeople() ([]models.Customer, error) {
	db, err := getConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM " + config.UsersTableName + " ORDER BY id")
	if err != nil {
		return []models.Customer{}, err
	}
	defer rows.Close()

	custs := make([]models.Customer, 0, 10)

	for rows.Next() {
		c := models.Customer{}
		err := rows.Scan(&c.Id, &c.Login, &c.Password, &c.Role, &c.Name, &c.Surname, &c.Age)
		if err != nil {
			log.Println(err)
			continue
		}

		custs = append(custs, c)
	}

	return custs, nil
}

func InsertCustomer(customer models.Customer) error {
	db, err := getConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO "+config.UsersTableName+"(login, password, role, name, surname, age) VALUES ($1, $2, $3, $4, $5, $6)",
		customer.Login, customer.Password, customer.Role, customer.Name, customer.Surname, customer.Age)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func UpdateCustomer(customer models.Customer) error {
	db, err := getConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE "+config.UsersTableName+" SET login=$1, password=$2, role=$3, name=$4, surname=$5, age=$6 WHERE id=$7;",
		customer.Login, customer.Password, customer.Role, customer.Name, customer.Surname, customer.Age, customer.Id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeleteCustomer(id int) error {
	db, err := getConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM "+config.UsersTableName+" WHERE id=$1", id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetPerson(id int) (models.Customer, error) {
	db, err := getConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	c := models.Customer{}
	row := db.QueryRow("SELECT * FROM "+config.UsersTableName+" WHERE id=$1", id)
	err = row.Scan(&c.Id, &c.Login, &c.Password, &c.Role, &c.Name, &c.Surname, &c.Age)
	if err != nil {
		return models.Customer{}, err
	}
	return c, nil
}

func GetPersonWithLoginAndPassword(login, password string) (models.Customer, error) {
	db, err := getConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	c := models.Customer{}

	row := db.QueryRow("SELECT * FROM "+config.UsersTableName+" WHERE login=$1 AND password=$2", login, password)
	err = row.Scan(&c.Id, &c.Login, &c.Password, &c.Role, &c.Name, &c.Surname, &c.Age)
	if err != nil {
		return models.Customer{}, err
	}
	return c, nil

}

package db

import (
	"User-service/models"
	"database/sql"
	"log"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) CreateNewUSer(req models.UserRequest) (*models.UserResponse, error) {
	query := `
		INSERT INTO USERS (name, email, password) 
		VALUES ($1,$2,$3)
		RETURNING name, email, password;
	`

	row := u.db.QueryRow(query, req.Name, req.Email, req.Password)

	var res models.UserResponse

	if err := row.Scan(&res.Name, &res.Email, &res.Password); err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

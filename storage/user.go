package storage

import (
	"database/sql"

	"app/models"
)

func Create(db *sql.DB, user models.User) (string, error) {

	var (
		id string
		query string
	)

	query = `
		INSERT INTO 
			users (first_name, last_name)
		VALUES ( $1, $2 )
		RETURNING id
	`
	err := db.QueryRow(
		query, 
		user.FirstName, 
		user.LastName,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetById(db *sql.DB, id string) (models.User, error) {

	var (
		user models.User
		query string
	)

	query = `
		SELECT
			id,
			first_name,
			last_name
		FROM
			users
		WHERE id = $1
	`
	err := db.QueryRow(
		query,
		id,
	).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetList(db *sql.DB) ([]models.User, error) {

	var (
		users []models.User
		query string
	)

	query = `
		SELECT
			id,
			first_name,
			last_name
		FROM
			users
	`
	rows, err := db.Query(query)

	if err != nil {
		return []models.User{}, err
	}

	for rows.Next() {
		var user models.User

		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
		)

		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}


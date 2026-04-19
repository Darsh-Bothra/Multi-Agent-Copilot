package repository

import "api/internal/db"




func CreateUser(name, email, password string) (string, error){
	query := `
		INSERT INTO user (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id string;

	err := db.DB.QueryRow(query, name, email, password).Scan(&id);

	if err != nil {
		return "", err;
	}

	return id, nil;
}

// it fetches the user
func GetUserByEmail(email string) (string, string, error) {
	query := `
		SELECT id, password
		FROM users
		WHERE email = $1
	`

	var id, password string;
	err := db.DB.QueryRow(query, email).Scan(&id, &password);

	if err != nil {
		return "", "", err;
	}

	return id, password, nil;
}
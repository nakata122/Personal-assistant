package users

import (
	"log"

	"server/internal/config"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context, userData *User) {
	// Insert into PostgreSQL
	var query string;

	switch userData.Role {
	case "user":
		query = `INSERT INTO users (google_id, email, name, picture, role) VALUES ($1, $2, $3, $4, $5)`;
	case "guest":
		query = `INSERT INTO users (google_id, email, name, picture, role) VALUES ($1, concat('temp', currval('users_user_id_seq'), '@tempmail.com'), $2, $3, $4)`;
	}

	query += `RETURNING user_id, email`;

	err := config.DbConn.QueryRow(c, query, userData.GoogleID, userData.Name, userData.Picture, userData.Role).Scan(&userData.UserID, &userData.Email);
	if err != nil {
		log.Println(err);
		return;
	}
}

func GetUserByID(c *gin.Context, id int) *User{
	var user User;

	query := `SELECT * FROM users WHERE user_id=$1`;
    err := config.DbConn.QueryRow(c, query, id).Scan(&user.UserID, &user.Email, &user.GoogleID, &user.Role, &user.Name, &user.Picture);
	if err != nil {
		log.Println(err);
		return nil;
	}
	
	return &user;
}

func GetUserByEmail(c *gin.Context, email string) *User{
	var user User;

	query := `SELECT * FROM users WHERE email=$1`;
    err := config.DbConn.QueryRow(c, query, email).Scan(&user.UserID, &user.Email, &user.GoogleID, &user.Role, &user.Name, &user.Picture);
	if err != nil {
		log.Println(err);
		return nil;
	}
	
	return &user;
}

func UpdateUser() {

}

func DeleteUser() {

}
package users

import (
	"log"

	"server/internal/config"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context, userData *User) int{
	// Insert into PostgreSQL
	var id int;
	var query string;

	switch userData.Role {
	case "user":
		query = `INSERT INTO users (google_id, email, name, picture, role) VALUES ($1, $2, $3, $4, $5)`;
	case "guest":
		query = `INSERT INTO users (google_id, email, name, picture, role) VALUES ($1, concat('temp', nextval('users_user_id_seq'), '@tempmail.com'), $2, $3, $4)`;
	}

	err := config.DbConn.QueryRow(c, query, userData.GoogleID, userData.Name, userData.Picture, userData.Role).Scan(&id);
	if err != nil {
		log.Println(err);
		return -1;
	}
    

	return id;
}

func GetUserByID() {

}

func GetUserByEmail(c *gin.Context, email string) *User{
	var user User;

	log.Println(email)
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
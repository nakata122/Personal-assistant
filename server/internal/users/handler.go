package users

import (
	"log"

	"server/internal/config"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context, userData *User) int{
	// Insert into PostgreSQL
	var id int;
    query := `INSERT INTO users (google_id, email, name, picture) VALUES ($1, $2, $3, $4)`;
    err := config.DbConn.QueryRow(c, query, userData.GoogleID, userData.Email, userData.Name, userData.Picture).Scan(&id);
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
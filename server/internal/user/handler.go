package user

import (
	"log"

	"server/internal/config"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context, userData *User) {
	// Insert into PostgreSQL
    query := `INSERT INTO users (google_id, email, name, picture) VALUES ($1, $2, $3, $4)`;
    _, err := config.DbConn.Exec(c, query, userData.GoogleID, userData.Email, userData.Name, userData.Picture);

    if err != nil {
		log.Println(err);
        return;
    }
}

func ReadUser() {

}

func UpdateUser() {

}

func DeleteUser() {

}
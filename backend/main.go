package main
import (
	"os"
	"fmt"
    "bufio"
    "github.com/gin-gonic/gin"
)

func ping(c *gin.Context) { 
	c.JSON(200, gin.H {
		"message": "pong",
	})
}

func getGoogleAuth(c *gin.Context) {
	fmt.Println(c.Request.Body);
}

func main() {
	router := gin.Default();
	router.GET("./api/auth/google", getGoogleAuth);
	router.GET("./api/ping", ping);

	
	env, err := os.Open(".env");
	if err != nil {
		fmt.Println(err)
	}
	
	s := bufio.NewScanner(env);
	s.Split(bufio.ScanWords);

	for s.Scan() {
		key := s.Text();
		s.Scan();

		os.Setenv(key, s.Text());
	}

	defer env.Close();

	port := os.Getenv("PORT");
	fmt.Println("Starting server on port: " + port);
	router.Run("localhost:" + port);
}
package utils

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func ParseEmail(c *gin.Context, message string) {
	cmd := exec.Command("python", "./scripts/script.py", message);

	out, err := cmd.CombinedOutput();
	if err != nil {
		fmt.Println("Error:", err);
	}

	fmt.Println(string(out));
}
package controllers

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func ExtractInfo(c *gin.Context, message string) {
	cmd := exec.Command("python", "./python/script.py", message);

	out, err := cmd.CombinedOutput();
	if err != nil {
		fmt.Println("Error:", err);
	}

	fmt.Println(string(out));
}
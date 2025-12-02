package controllers

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func ExtractInfo(c *gin.Context) {
	cmd := exec.Command("python", "./python/script.py", "hello this is an email pls I am Naiden Kostov.");

	out, err := cmd.CombinedOutput();
	if err != nil {
		fmt.Println("Error:", err);
	}

	fmt.Println(string(out));
}
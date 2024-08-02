package main

import (
	"database/sql"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "host=db user=postgres password=yourpassword dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	r.POST("/api/strong_password_steps", strongPasswordSteps)
	r.Run(":8080")
}

func strongPasswordSteps(c *gin.Context) {
	var request struct {
		InitPassword string `json:"init_password"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	steps := calculateSteps(request.InitPassword)

	// Log request and response
	_, err := db.Exec("INSERT INTO logs (request, response) VALUES ($1, $2)", request.InitPassword, steps)
	if err != nil {
		log.Printf("Error logging to database: %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"num_of_steps": steps})
}

func calculateSteps(password string) int {
	// Implementation of password strength calculation
	// (This is a simplified version, you may want to expand on this)
	steps := 0
	if len(password) < 6 {
		steps += 6 - len(password)
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		steps++
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		steps++
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		steps++
	}
	return steps
}

package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func InitDB() {
	filePath := "./database/db.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open database file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		Users = append(Users, CreateUserFromLine(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	log.Printf("Users scanned successfully. Total: %d", len(Users))
}

func CreateUserFromLine(l string) User {
	var u User
	data := strings.Split(l, " ")
	u.Id = data[0]
	u.Username = data[1]
	u.Password = data[2]
	return u
}

func FindUserFromUsers(l, p string) bool {
	for _, u := range Users {
		if l == u.Username && p == u.Password {
			return true
		}
	}
	return false
}

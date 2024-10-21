package database

import (
    "golang.org/x/crypto/bcrypt"
    "fmt"
)

func Seed() {
    adminPassword := "adminpassword"
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)

    admin := User{Username: "admin", Password: string(hashedPassword), IsAdmin: true}
    DB.Create(&admin)

    fmt.Println("Seeded Admin User:", admin.Username, "Hashed Password:", admin.Password) // Debug line
}


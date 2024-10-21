package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

var DB *gorm.DB

func InitDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    DB.AutoMigrate(&User{})
}

type User struct {
    gorm.Model
    Username string
    Password string
    IsAdmin  bool
}

func HashPassword(password string) string {
    hash := sha256.New()
    hash.Write([]byte(password))
    hashedPassword := hex.EncodeToString(hash.Sum(nil))
    fmt.Println("Hashed Password:", hashedPassword) // Debug line
    return hashedPassword
}


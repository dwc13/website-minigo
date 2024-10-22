package database

import (
    "crypto/sha256"
    "encoding/hex"
    "log"
    "os"
    "website-minigo/utils"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDatabase() {
    err := godotenv.Load() // load .env file
    if err != nil {
        log.Fatalf("failed to load env: %v", err)
    }
    secretKey := os.Getenv("DB_SECRET_KEY")
    if secretKey == "" {
        log.Fatal("DB_SECRET_KEY environment variable not set")
    }

    if _, err := os.Stat("encrypted.db"); os.IsNotExist(err) {
        log.Println("encrypted.db does not exist. Creating initial unencrypted database.")
        DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
        if err != nil {
            log.Fatalf("failed to connect database: %v", err)
        }
        DB.AutoMigrate(&User{})
        log.Println("Initial unencrypted database created.")

        data, err := os.ReadFile("users.db")
        if err != nil {
            log.Fatalf("failed to read initial database: %v", err)
        }
        err = utils.EncryptFile("encrypted.db", data, secretKey)
        if err != nil {
            log.Fatalf("failed to encrypt initial database: %v", err)
        }
        os.Remove("users.db")
        log.Println("Initial unencrypted database encrypted and removed.")
    }

    log.Println("Decrypting database.")
    decryptedDB, err := utils.DecryptFile("encrypted.db", secretKey)
    if err != nil {
        log.Fatalf("failed to decrypt database: %v", err)
    }
    err = os.WriteFile("users.db", decryptedDB, 0644)
    if err != nil {
        log.Fatalf("failed to write decrypted database: %v", err)
    }
    log.Println("Database decrypted and written to users.db.")

    DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    DB.AutoMigrate(&User{})
    log.Println("Database connected and migrations run.")
}

func CloseDatabase() {
    secretKey := os.Getenv("DB_SECRET_KEY")
    if secretKey == "" {
        log.Fatal("DB_SECRET_KEY environment variable not set")
    }

    log.Println("Closing database and encrypting users.db.")
    data, err := os.ReadFile("users.db")
    if err != nil {
        log.Fatalf("failed to read database: %v", err)
    }
    err = utils.EncryptFile("encrypted.db", data, secretKey)
    if err != nil {
        log.Fatalf("failed to encrypt database: %v", err)
    }
    os.Remove("users.db")
    log.Println("Database users.db encrypted and removed.")
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
    return hex.EncodeToString(hash.Sum(nil))
}


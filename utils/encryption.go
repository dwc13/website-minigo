package utils

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
    "os"
)

func EncryptFile(filename string, data []byte, key string) error {
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return err
    }

    ciphertext := gcm.Seal(nonce, nonce, data, nil)
    return os.WriteFile(filename, ciphertext, 0644)
}

func DecryptFile(filename string, key string) ([]byte, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    nonceSize := gcm.NonceSize()
    if len(data) < nonceSize {
        return nil, io.ErrUnexpectedEOF
    }

    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    return gcm.Open(nil, nonce, ciphertext, nil)
}


package hashdemo

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func calculateHash(toBeHash string) (hashed string) {
	hashedBytes := sha256.Sum256([]byte(toBeHash))
	hashed = hex.EncodeToString(hashedBytes[:])
	log.Printf("toBeHash %s hashed %s", toBeHash, hashed)
	return
}

func CalculateHash(toBeHash string) (hashed string) {
	return calculateHash(toBeHash)
}

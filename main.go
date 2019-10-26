package main

import (
	"log"

	"github.com/corganfuzz/stringencrypter/utils"
)

func main() {
	key := "111023043350789514532147"
	message := "I am a message"
	log.Println("Original Message: ", message)

	encryptedString := utils.EncryptedString(key, message)
	log.Println("Encrypterd message: ", encryptedString)

	decryptedString := utils.DecryptedString(key, encryptedString)
	log.Println("Decrypted message: ", decryptedString)

}

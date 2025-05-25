package shortener

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))

	return algo.Sum(nil)
}

func Base58Encode(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		log.Printf("Failed to encode msg: %v", err)
		os.Exit(1)
	}

	return string(encoded)
}

func GenerateShortLink(initialLink, userID string) string {
	urlHashBytes := sha256Of(initialLink + userID)
	generatedNum := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := Base58Encode([]byte(fmt.Sprintf("%d", generatedNum)))

	return finalString[:8]
}

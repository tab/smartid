package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/tab/smartid/internal/errors"
)

const (
	// HashTypeSHA256 is the hash type for SHA-256
	HashTypeSHA256 = "SHA256"

	// HashTypeSHA384 is the hash type for SHA-384
	HashTypeSHA384 = "SHA384"

	// HashTypeSHA512 is the hash type for SHA-512
	HashTypeSHA512 = "SHA512"

	// RandomBytesLength is length of random bytes
	RandomBytesLength = 64
)

// GenerateHash generates a random hash based on the given hash type
func GenerateHash(hashType string) (string, error) {
	randBytes := make([]byte, RandomBytesLength)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", errors.ErrFailedToGenerateRandomBytes
	}

	var encodedHash string

	switch strings.ToUpper(hashType) {
	case HashTypeSHA256:
		hash := sha256.Sum256(randBytes)
		encodedHash = base64.StdEncoding.EncodeToString(hash[:])
	case HashTypeSHA384:
		hash := sha512.Sum384(randBytes)
		encodedHash = base64.StdEncoding.EncodeToString(hash[:])
	case HashTypeSHA512:
		hash := sha512.Sum512(randBytes)
		encodedHash = base64.StdEncoding.EncodeToString(hash[:])
	default:
		return "", errors.ErrUnsupportedHashType
	}

	return encodedHash, nil
}

// GenerateVerificationCode generates a verification code based on the given hash
func GenerateVerificationCode(hash string) (string, error) {
	decodedHash, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return "", err
	}

	sha256Hash := sha256.Sum256(decodedHash)
	lastTwoBytes := sha256Hash[len(sha256Hash)-2:]
	codeInt := binary.BigEndian.Uint16(lastTwoBytes)
	vc := codeInt % 10000
	code := fmt.Sprintf("%04d", vc)

	return code, nil
}

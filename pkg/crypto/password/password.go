package password

import "golang.org/x/crypto/bcrypt"

func Generate(s string) (string, error) {
	inputBytes := []byte(s)
	encryptedBytes, err := bcrypt.GenerateFromPassword(inputBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	encryptedHash := string(encryptedBytes)
	return encryptedHash, nil
}

func Compare(hash string, s string) error {
	providedPassword := []byte(s)
	storedHash := []byte(hash)
	return bcrypt.CompareHashAndPassword(storedHash, providedPassword)
}

package util

import (
	"github.com/satori/go.uuid"
	"log"
)

func GenerateUUIDV4() (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		log.Printf("GenerateUUIDV4 went wrong: %s", err)
		return "", err
	}

	return uid.String(), nil
}

package util

import (
	"errors"
	"github.com/satori/go.uuid"
	"log"
	"strings"
)

func GenerateUUIDV4() (string, error) {
	uid := uuid.NewV4().String()
	if len(uid) == 0 || len(strings.TrimSpace(uid)) == 0 {
		log.Print("GenerateUUIDV4 went wrong.")
		return "", errors.New("GenerateUUIDV4 went wrong")
	}

	return uid, nil
}

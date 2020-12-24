package utils

import uuid "github.com/satori/go.uuid"

func RandomUUID() (uuid.UUID, error) {
	return uuid.NewV4(), nil
}

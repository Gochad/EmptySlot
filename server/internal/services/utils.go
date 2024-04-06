package services

import (
	"strconv"

	"github.com/google/uuid"
)

func generateUUID() string {
	return strconv.Itoa(int(uuid.New().ID()))
}

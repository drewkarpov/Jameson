package pkg

import (
	"encoding/hex"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func GetNewId() string {
	id, _ := uuid.New()
	return hex.EncodeToString(id[:])
}

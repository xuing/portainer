package database

import (
	"fmt"

	portainer "github.com/portainer/portainer/api"
	"github.com/portainer/portainer/api/database/boltdb"
)

// NewDatabase should use config options to return a connection to the requested database
func NewDatabase(storeType, storePath, encryptionKey string) (connection portainer.Connection, err error) {
	switch storeType {
	case "boltdb":
		isDBEncrypted := encryptionKey != ""
		return &boltdb.DbConnection{
			Path:          storePath,
			EncryptionKey: encryptionKey,
			IsDBEncrypted: isDBEncrypted,
		}, nil
	}
	return nil, fmt.Errorf("unknown storage database: %s", storeType)
}

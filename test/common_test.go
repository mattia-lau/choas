package test

import (
	"chaos-go/common"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectingDatabase(t *testing.T) {
	dbPath := "test.db"

	asserts := assert.New(t)
	db, _ := common.Init(dbPath).DB()
	// Test create & close DB
	_, err := os.Stat(dbPath)
	asserts.NoError(err, "Db should exist")
	asserts.NoError(db.Ping(), "Db should be able to ping")

	// Test get a connecting from connection pools
	connection := common.GetDB()
	asserts.NoError(connection.Error, "Db should be able to ping")
	defer db.Close()

	// Test DB exceptions
	os.Chmod(dbPath, 0000)
	db, _ = common.Init(dbPath).DB()
	asserts.Error(db.Ping(), "Db should not be able to ping")
	db.Close()
	os.Chmod(dbPath, 0644)
}

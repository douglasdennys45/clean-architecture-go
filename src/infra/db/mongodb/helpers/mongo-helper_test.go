package helpers

import (
	"fmt"
	"github.com/benweissmann/memongo"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestConnect(t *testing.T) {
	mongoServer, err := memongo.Start("4.0.5")
	if err != nil {
		log.Fatal(err)
	}
	defer mongoServer.Stop()

	err = Connect(mongoServer.URI(), memongo.RandomDatabase())
	assert.Nil(t, err)

	fmt.Println(mongoServer)

	collection := GetCollection("users")
	assert.NotNil(t, collection)

	err = Disconnect()
	assert.Nil(t, err)
}

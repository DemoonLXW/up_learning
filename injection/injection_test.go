package injection

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvideDataBase(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	client, err := ProvideDataBase()
	defer client.Close()
	assert.Nil(t, err)
}

package injection

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvideDataBase(t *testing.T) {
	os.Setenv("DB_CONFIG", "../database.config.json")
	db, err := ProvideDataBase()
	assert.Nil(t, err)
	assert.Equal(t, true, db.Config.SkipDefaultTransaction)
}

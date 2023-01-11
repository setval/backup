package config_test

import (
	"testing"

	"github.com/setval/container-backup/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	cfg, err := config.Parse("../../config.test.yml")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, cfg.Storage.Type, config.TypeStorageS3)
	assert.Equal(t, cfg.Storage.S3Config.Endpoint, "endpoint")
	assert.Equal(t, cfg.Storage.S3Config.AccessKey, "access_key")
	assert.Equal(t, cfg.Storage.S3Config.SecretKey, "secret_key")

	assert.Equal(t, 2, len(cfg.Databases))
	assert.Equal(t, "testdb1", cfg.Databases[0].Name)
	assert.Equal(t, "testdb2", cfg.Databases[1].Name)
	assert.Equal(t, 2, len(cfg.Databases[0].Tables))
	assert.Equal(t, "table1", cfg.Databases[0].Tables[0])
	assert.Equal(t, "table2", cfg.Databases[0].Tables[1])
}

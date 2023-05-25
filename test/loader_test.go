package test

import (
	"github.com/go-for-go/yak/src/services/ConfigLoader"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestRegisterConfig(t *testing.T) {
	os.Setenv("TEST_PJ_PG_DEFAULT_USER", "set_via_env")
	err := ConfigLoader.RegisterConfig("./etc/", func() error { return nil }, "test.toml")

	require.NoError(t, err)

	// get param from common.toml
	require.Equal(t, "my-db", viper.GetString("pg.default.database"))

	// overwrite by test.toml
	require.Equal(t, "test:1234", viper.GetString("pg.default.host"))

	// overwrite by ENV
	require.Equal(t, "set_via_env", viper.GetString("pg.default.user"))
}

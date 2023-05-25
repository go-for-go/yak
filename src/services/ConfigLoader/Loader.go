package ConfigLoader

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

func RegisterConfig(path string, validate func() error, files ...string) error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}
	exPath := filepath.Dir(ex)

	viper.AddConfigPath(path)
	viper.AddConfigPath(exPath)
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	viper.SetConfigName("common")
	viper.ReadInConfig()

	for _, file := range files {
		if file != "" {
			dir := filepath.Dir(file)
			if dir == "." {
				file = path + file
			}

			var _, err = os.Stat(file)
			if !os.IsNotExist(err) {
				viper.SetConfigFile(file)
				viper.MergeInConfig()
			} else {
				fmt.Printf("Config file \"%s\" was not found\n", file)
			}
		}
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetEnvPrefix(viper.GetString("module"))

	viper.AutomaticEnv()

	if err := validate(); err != nil {
		return fmt.Errorf("config is not valid : %s", err.Error())
	}
	return nil
}

func ReadConfigFilesV3(path, file, prefix string, validate func() error) {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	viper.AddConfigPath(path)
	viper.AddConfigPath(exPath)
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	viper.SetConfigName("common")
	viper.ReadInConfig()

	if file != "" {
		file = path + file
		var _, err = os.Stat(file)
		if !os.IsNotExist(err) {
			viper.SetConfigFile(file)
			viper.MergeInConfig()
		} else {
			panic(fmt.Sprintf("Config file \"%s\" was not found", file))
		}
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetEnvPrefix(prefix)

	viper.AutomaticEnv()

	err = validate()
	if err != nil {
		panic(fmt.Sprintf("config is not valid : %s", err.Error()))
	}
}

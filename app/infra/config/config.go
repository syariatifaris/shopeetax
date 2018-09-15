package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/syariatifaris/shopeetax/app/infra/config/toml"
)

//Database configuration
type Database struct {
	Host     string
	Type     string
	Name     string
	User     string
	Password string
	Port     int64
}

type DbConfig struct {
	Database Database
}

type ConfigurationData struct {
	Database Database
}

var (
	directories = []string{
		"files/conf/",
		"../../files/conf/",
		"../../../files/conf/",
		"../src/github.com/syariatifaris/bannerapp/files/conf/",
	}

	extension = "toml"
)

type Config interface {
	DecodeConfig(c interface{}) error
}

//NewConfiguration reads all configuration
func NewConfiguration() *ConfigurationData {
	var dbCfg DbConfig

	err := readConfigurationModule("db", extension, directories, &dbCfg)
	if err != nil {
		panic(fmt.Sprintf("unable to resolve db configuration with err %s", err.Error()))
	}

	return &ConfigurationData{
		Database: dbCfg.Database,
	}
}

const (
	ExtTomlFile = "toml"
)

func readConfigurationModule(moduleName, ext string, dirs []string, cfgContainer interface{}) error {
	filePath := resolveConfigModuleFileLoc(moduleName, ext, dirs)
	if filePath == "" {
		return errors.New("unable to find files path")
	}

	switch ext {
	case ExtTomlFile:
		configuration := toml.NewTomlConfiguration(filePath)
		err := configuration.DecodeConfig(cfgContainer)
		if err != nil {
			return err
		}
	default:
		return errors.New("unrecognized config file extension")
	}

	return nil
}

func resolveConfigModuleFileLoc(module, ext string, dirs []string) string {
	for _, dir := range dirs {
		file := fmt.Sprintf("./%s%s.%s", dir, module, ext)
		if _, err := os.Stat(file); !os.IsNotExist(err) {
			return file
		}
	}
	return ""
}

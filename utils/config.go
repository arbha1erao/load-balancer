package utils

import (
	"encoding/json"
	"os"
)

func LoadConfig(configPath string, config interface{}) error {
	configFile, err := os.Open(configPath)
	if err != nil {
		Logger.Error().Err(err).Msg("[CONFIG-LOADER] failed to open config file")
		return err
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(config)
	if err != nil {
		Logger.Error().Err(err).Msg("[CONFIG-LOADER] failed to decode config")
		return err
	}

	return nil
}

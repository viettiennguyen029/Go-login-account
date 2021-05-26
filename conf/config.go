package conf

import (
	"encoding/json"
	"log"
	"os"
)

//ConfigOption contains config of the server
type ConfigOption struct {
	Addr string
}

var (
	config *ConfigOption
)

//GetConfigOption returns the ConfigOption
func GetConfigOption() *ConfigOption {
	if config != nil {
		return config
	}
	file, err := os.Open("./conf/config.json")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return config
}
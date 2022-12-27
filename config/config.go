package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

// gloabl constants
var (
	configPath     = "."         // path where the configuration file is stored
	configFileName = "name.json" // configuration file name
)

type GlobalConfig struct {
	// state the json based key/value pair
	// format Name DataType `json:"json_name"`
	SampleKey           string           `json:"sampleKey"`
	SampleKeyWithObject SampleWithObject `json:"sampleKeyWithObject"`
}

type SampleWithObject struct {
	Key1 string `json:"key_1"`
}

// global config storage variable
var config GlobalConfig

// a sync to only load to once throughout the life of the program
var configOnce sync.Once

// helper function to updated the path with respect to the path of running the code
func SetConfFilePath(path string) {
	configPath = path
}

// helper function to update the configurations file name
func SetConfFileName(name string) {
	configFileName = name
}

// Load and read the configurations file once and store it in config variable
func GetConfig() *GlobalConfig {
	configOnce.Do(func() {
		// this function is set to execute just once through out thew life of the program
		// will read the file in forms on bytes
		bytes, err := ioutil.ReadFile(configPath + "/" + configFileName)
		if err != nil {
			// handle the error, i.e. load some default configs
			// panic in proiduction is not a good option
			// panic(err)
			config = *defaultConfigurations()
			return
		}

		// will unmarshal or decode the bytes and store it in the config global variable
		// matching the keys to the `json:"json_name"` tag
		err = json.Unmarshal(bytes, &config)
		if err != nil {
			panic(err)
		}
	})
	return &config
}

func defaultConfigurations() *GlobalConfig {
	// define the default state if, e.g. configuration file is not defined
	return &GlobalConfig{}
}

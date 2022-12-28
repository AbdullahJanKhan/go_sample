package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"
)

// gloabl constants
var (
	configPath     = "."           // path where the configuration file is stored
	configFileName = "config.json" // configuration file name
)

type GlobalConfig struct {
	// state the json based key/value pair
	// format Name DataType `json:"json_name"`
	SampleKey           string           `json:"sampleKey"`
	SampleKeyWithObject SampleWithObject `json:"sampleKeyWithObject"`
	Aragon              Aragon           `json:"aragon"`
	DataSource          DataSourceConfig `json:"dataSource"`
	Grpc                GrpcConfig       `json:"grpc"`
	JwtSecret           string           `json:"jwtSecret"`
	Kafka               KafkaConfig      `json:"kafka"`
	Redis               RedisConfig      `json:"redis"`
	RestServer          RestServerConfig `json:"restServer"`
	ExternalGrpc        ExternalGrpc     `json:"externalGrpc"`
	Logger              LoggerConfig     `json:"logger"`
}

type RestServerConfig struct {
	Addr string `json:"addr"`
}

type DataSourceConfig struct {
	DriverName        string `json:"driverName"`
	Addr              string `json:"addr"`
	Port              string `json:"port"`
	Database          string `json:"database"`
	User              string `json:"user"`
	Password          string `json:"password"`
	EnableAutoMigrate bool   `json:"enableAutoMigrate"`
	Retries           int64  `json:"retries"`
}

type GrpcConfig struct {
	Port          string        `json:"port"`
	Addr          string        `json:"addr"`
	ServerMinTime time.Duration `json:"serverMinTime" default:"5m"` // if a client pings more than once every 5 minutes (default), terminate the connection
	ServerTime    time.Duration `json:"serverTime" default:"2h"`
	ServerTimeOut time.Duration `json:"serverTimeOut" default:"20s"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

type KafkaConfig struct {
	Brokers []string `json:"brokers"`
}

type Aragon struct {
	AragonPepper  string `json:"aragonPepper"`
	AragonTime    uint32 `json:"aragonTime"`
	AragonMemory  uint32 `json:"aragonMemory"`
	AragonKeyLen  uint32 `json:"aragonKeyLen"`
	AragonThreads uint8  `json:"aragonThreads"`
	AragonSaltLen uint32 `json:"aragonSaltLen"`
}

type ExternalGrpc struct {
	Sample GrpcConn `json:"sampleGrpc"`
}

type GrpcConn struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
}

type LoggerConfig struct {
	LogLevel       string `json:"logLevel"`
	LogEnvironment string `json:"logEnvironment"`
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

package conf

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Configuration struct {
	Endpoint struct {
		Ethereum struct {
			Mainnet string
			Ropsten string
		}
	}
}

var once sync.Once

var configuration = &Configuration{}

func init() {
	if os.Getenv("GOPATH") == "" {
		panic("GOPATH is not set!")
	}
}

func GetConfiguration() *Configuration {
	once.Do(func() {
		viper.AutomaticEnv()
		viper.SetConfigFile(os.Getenv("GOPATH") + "/src/github.com/DE-labtory/zulu/conf/config.yaml")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
		err := viper.Unmarshal(&configuration)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
	})
	return configuration
}

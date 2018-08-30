package conf

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

//Vars contain all configuration variables
type Vars struct {
	App string
	Log struct {
		Error string
	}
	Port    string
	LevelDB struct {
		Path string
	}
}

//Read reads all the environment variables using the envconfig package
func Read(name string) *Vars {
	vars := &Vars{
		App: name,
	}
	v := viper.New()
	v.SetConfigName("conf")
	v.AddConfigPath("./src/conf")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		os.Exit(1)
	}
	if err := v.Unmarshal(&vars); err != nil {
		log.Fatalf("Error unmarshalling config file, %s", err)
		os.Exit(1)
	}

	return vars
}

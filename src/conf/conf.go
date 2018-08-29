package conf

import (
	"log"

	"github.com/spf13/viper"
)

//Vars contain all configuration variables
type Vars struct {
	App string
	Log struct {
		Error string
	}
	Port string
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
	}
	if err := v.Unmarshal(&vars); err != nil {
		log.Fatalf("Error unmarshalling config file, %s", err)
	}

	return vars
}

package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var readFromLocalAppYml = func(v *viper.Viper) *viper.Viper {
	fmt.Println("Reading application yml configuration")
	v.SetConfigName(defaultConfigFileName)
	v.SetTypeByDefaultValue(true)
	v.SetConfigType(defaultConfigFileType)
	v.AddConfigPath("./" + defaultConfigPath)
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Viper read config has an error; %e\n", err)
	}
	return v
}

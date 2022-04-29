package config

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

var readFromEnv = func(v *viper.Viper) *viper.Viper {
	fmt.Println("Reading environment configuration")
	addKeysToViper(v)
	v.AutomaticEnv()
	return v
}

func addKeysToViper(v *viper.Viper) {
	var reply interface{} = Config{}
	t := reflect.TypeOf(reply)
	keys := getAllKeys(t)
	for _, key := range keys {
		v.SetDefault(key, "")
	}
}

func getAllKeys(t reflect.Type) []string {
	var result []string
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		n := strings.ToUpper(f.Tag.Get(ymlTagName))
		if reflect.Struct == f.Type.Kind() {
			subkeys := getAllKeys(f.Type)
			for _, k := range subkeys {
				result = append(result, n+"."+k)
			}
		} else {
			result = append(result, n)
		}
	}
	return result
}

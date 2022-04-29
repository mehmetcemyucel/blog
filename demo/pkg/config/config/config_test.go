package config

import (
	"github.com/spf13/viper"
	"testing"
)

func TestReadConfig_whenEnvEmpty_thenReadLocalFile(t *testing.T) {
	expectedValue := ":8080"
	readFromLocalAppYml = func(v *viper.Viper) *viper.Viper {
		v.SetDefault("APP-SERVER.PORT", expectedValue)
		return v
	}
	c := ReadConfig(&Config{}, "")
	if c.Server.Port != expectedValue {
		t.Errorf("Expected server port variable %v actual %v", expectedValue, c.Server.Port)
	}
}

func TestReadConfig_whenEnv_thenReadFromEnv(t *testing.T) {
	expectedValue := ":8080"
	readFromEnv = func(v *viper.Viper) *viper.Viper {
		v.SetDefault("APP-SERVER.PORT", expectedValue)
		return v
	}
	c := ReadConfig(&Config{}, "ENV")
	if c.Server.Port != expectedValue {
		t.Errorf("Expected server port variable %v actual %v", expectedValue, c.Server.Port)
	}
}

func TestReadConfig_whenEnvRemote_thenReadFromRemoteConfigServer(t *testing.T) {
	expectedValue := ":8080"
	readFromConfigServer = func(v *viper.Viper) *viper.Viper {
		v.SetDefault("APP-SERVER.PORT", expectedValue)
		return v
	}
	c := ReadConfig(&Config{}, "REMOTE")
	if c.Server.Port != expectedValue {
		t.Errorf("Expected server port variable %v actual %v", expectedValue, c.Server.Port)
	}
}

func TestReadConfig_whenEnvLocal_thenReadLocalFile(t *testing.T) {
	expectedValue := ":8080"
	readFromLocalAppYml = func(v *viper.Viper) *viper.Viper {
		v.SetDefault("APP-SERVER.PORT", expectedValue)
		return v
	}
	c := ReadConfig(&Config{}, "LOCAL")
	if c.Server.Port != expectedValue {
		t.Errorf("Expected server port variable %v actual %v", expectedValue, c.Server.Port)
	}
}

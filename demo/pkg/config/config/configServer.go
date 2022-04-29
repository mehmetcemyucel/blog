package config

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"os"
)

var readFromConfigServer = func(v *viper.Viper) *viper.Viper {
	rca, rcu, rcp := readEnvVar()
	fmt.Printf("Reading config server configuration; url: %s\n", rca)
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.Set(fasthttp.HeaderAuthorization, "Basic "+base64.StdEncoding.EncodeToString([]byte(rcu+":"+rcp)))
	req.SetRequestURI(rca)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Printf("Config server read client get failed: %s\n", err.Error())
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		fmt.Printf("Expected status code %d but got %d\n", fasthttp.StatusOK, resp.StatusCode())
	} else {
		v.SetConfigType(defaultConfigFileType)
		if err := v.ReadConfig(bytes.NewBuffer(resp.Body())); err != nil {
			fmt.Printf("Viper read config has an error; %e\n", err)
		}
	}
	return v
}

func readEnvVar() (string, string, string) {
	rca, readed := os.LookupEnv("REMOTE_CONFIG_ADDRESS")
	if !readed {
		fmt.Println("REMOTE_CONFIG_ADDRESS should be given")
	}
	rcu, readed := os.LookupEnv("REMOTE_CONFIG_USERNAME")
	if !readed {
		fmt.Println("REMOTE_CONFIG_USERNAME should be given")
	}
	rcp, readed := os.LookupEnv("REMOTE_CONFIG_PASSWORD")
	if !readed {
		fmt.Println("REMOTE_CONFIG_PASSWORD should be given")
	}
	return rca, rcu, rcp
}

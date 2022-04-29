package util

import (
	"fmt"
	"os"
	"strings"
)

const environmentKey = "ACTIVE_PROFILE"

var ReadEnv = func() string {
	p, _ := os.LookupEnv(environmentKey)
	fmt.Println("Active Profile: [" + p + "] readed from runtime arguments[" + environmentKey + "].")
	return strings.ToUpper(p)
}

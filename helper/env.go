package helper

import (
	"os"
)

const ENV_DEV = "dev"
const ENV_TEST = "test"
const ENV_PRODUCT = "product"

const LOCATION_LOCAL = "local"

func GetEnv() (env string) {
	env = os.Getenv("MACHINE_ENV")
	if env != "" {
		return env
	} else {
		return ENV_DEV
	}
}

func GetHostName() string {
	hostname, err := os.Hostname()
	// not has hostname? how
	if err != nil {
		// todo set machine id
		return ""
	}
	return hostname
}

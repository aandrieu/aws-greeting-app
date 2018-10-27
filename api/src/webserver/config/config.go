package config

import (
	"fmt"
	"os"
)

var METADATASTORE_URL = safeGetEnv("METADATASTORE_URL")

func safeGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if ok == true {
		return val
	} else {
		panic(fmt.Sprintf("ENV variable not found for key: %s", key))
	}
}

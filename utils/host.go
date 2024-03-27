package utils

import "os"

func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		name = "localhost"
	}
	return name
}

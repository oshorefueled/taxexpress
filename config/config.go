package config

import "strconv"

var (
	AdminKey = "123456789"
	Port = 3000
)

func PortToServe () string {
	return ":" + strconv.Itoa(Port)
}
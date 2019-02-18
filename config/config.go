package config

import "strconv"

var Port = 3000

func PortToServe () string {
	return ":" + strconv.Itoa(Port)
}
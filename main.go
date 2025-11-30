package main

import (
	"fmt"
	"redigo/config"
)

func main() {
	const configFile string = "./redis.conf"
	config.SetupConfig(configFile)
	fmt.Printf("%+v", config.Properties)
}

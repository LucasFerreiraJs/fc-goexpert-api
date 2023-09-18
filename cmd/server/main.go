package main

import "github.com/lucasferreirajs/fc-goexpert/apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	print(config.DBDriver)
}

package main

import "github.com/anton-ag/javacode/internal/app"

const configPath = "configs/config.yml"

func main() {
	app.Run(configPath)
}

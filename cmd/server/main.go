package main

import "github.com/silastgoes/CVGeneretor/configs"

func main() {
	config := configs.Load("./")
	println(config.Env.DBDriver)
}

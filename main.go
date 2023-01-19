package main

import (
	"financialproduct/config"
	"financialproduct/routers"
)

func main() {
	config.Init()
	r := routers.Init()

	r.Run("0.0.0.0:80")
}

package main

import (
	"github.com/rosas/onex/cmd/onex-nightwatch/app"
)

func main() {
	app.NewJobServer().Run()
}

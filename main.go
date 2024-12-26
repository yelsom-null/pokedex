package main

import (
	"poke/internal"
)




func main() {

	clientPoke := internal.NewClient()

	cfg := &config{
		client: clientPoke,
	}
	start(cfg)

}


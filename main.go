package main

import (
	"api-cli-golang/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	err := aplicacao.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

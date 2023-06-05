package main

import (
	"log"
	"os"
	"github.com/urfave/cli"
)

func main() {
	log.Printf("xxx")
	(&cli.App{}).Run(os.Args)
}

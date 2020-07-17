package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
)

var appEnv = flag.String("env", "prod", "Put the environment in production or in dev to load different configuration.")

func init() {
	flag.Parse()
	if *appEnv == "dev" {
		err := godotenv.Load(".env.dev.local", ".env")
		if err != nil {
			log.Fatal("Error loading .env, .env.dev.local file")
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	serve()
}

package main

import (
	"log"
	"os"
	"time"
)

const ENV_VAR_NAME = "hello"

func main() {
	log.Println("starting container!")
	for {
		// prints environment variable "hello" and its value
		log.Printf("%v: %v", ENV_VAR_NAME, os.Getenv(ENV_VAR_NAME))
		// sleeps for one second
		time.Sleep(time.Second * 1)
	}
}

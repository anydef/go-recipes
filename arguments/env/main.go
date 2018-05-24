package main

import (
	"os"
	"log"
	"fmt"
)

func main() {
	key := "DB"
	value, is_set := os.LookupEnv(key)
	if !is_set {
		log.Printf("Env variable is not set %s.\n", key)
	}
	fmt.Printf(value)
}

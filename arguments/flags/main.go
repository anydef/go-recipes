package main

import (
	"fmt"
	"flag"
	"strings"
	"log"
	"os"
)

type ArrayValue []string

func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

func (a *ArrayValue) Set(s string) error {
	*a = strings.Split(s, ",")
	return nil
}

func main() {
	retry := flag.Int("retry", -1, "")

	var logPrefix string
	flag.StringVar(&logPrefix, "prefix", "", "LoggerPrefix")

	var arr ArrayValue

	flag.Var(&arr, "array", "Input array iterate through")

	flag.Parse()

	logger := log.New(os.Stdout, logPrefix+" ", log.Ldate)

	retryCount := 0

	for retryCount < *retry {
		logger.Println("Retrying connection")
		logger.Printf("Sending array %v\n", arr)
		retryCount++
	}

}

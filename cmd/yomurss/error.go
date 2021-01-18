package main

import (
	"log"
	"os"
)

func check(err error, code int) {
	if err != nil {
		log.Println(err)
		os.Exit(code)
	}
}

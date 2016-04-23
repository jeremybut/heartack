package main

import (
	"log"
	"os"
)

var Log *log.Logger

func init() {
	Log = log.New(os.Stdout, "INFO:", log.LstdFlags)
}

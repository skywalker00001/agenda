package entity

import (
	"io"
	"log"
	"os"
)

// NewLogger create a logger which write info on screen and in ./data/log.txt with specific prefix
func NewLogger(prefix string) *log.Logger {
	file, _ := os.OpenFile("./data/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return log.New(io.MultiWriter(file, os.Stdout), prefix, log.Ldate|log.Ltime)
}

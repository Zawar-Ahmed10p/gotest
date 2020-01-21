package main

import (
	"github.com/Zawar-Ahmed10p/gotest/gologger"
)

func main() {
	logger := gologger.GetLogger()
	logger.Warning("9999999")
	for i := 0; i < 99999999; i++ {
		logger.Print("fsdfsdf")
		logger.Warning("9999999")
		// panic("das")
	}

}

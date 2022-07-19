package Log

import (
	"fmt"
	"io"
	"log"
	"os"
)

func InitLog() {
	logfile, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Print("logfile open err")
		StopLog(logfile)
		return
	}

	multiWriter := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiWriter)

	log.Println("Log module has initialized.")
}

func StopLog(f *os.File) {
	f.Close()
}

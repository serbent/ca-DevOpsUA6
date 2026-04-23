package main

import (
	"log"
	"os"
)

const (
	Ldate = iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	LstdFlags = Ldate | Ltime
)

func main() {

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	multi := log.MultiWriter(os.Stdout, file)
	log.SetOutput(multi) // os.Stdout is the default output for the standard logger multi is for writing to multiple outputs
	log.Println("This is our first logged message")
	log.SetFlags(log.Ldate) // print the date in the local time zone: 2009/01/23
	log.Println("This is our second logged message")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds) // print the date and time with microsecond precision: 2009/01/23 01:23:23.123123
	log.Println("This is our third logged message")
	log.SetFlags(log.Llongfile) // print the full path of the file and the line number: /home/user/go/src/main.go:10
	log.Println("This is our fourth logged message")
	log.SetFlags(log.Lshortfile) // print the file name and the line number: main.go:10
	log.Println("This is our fifth logged message")
	log.SetFlags(log.LUTC | log.Ldate | log.Ltime) // print the date and time in UTC: 2009/01/23 01:23:23
	log.Println("This is our sixth logged message")

	log.SetPrefix("INFO: ")

	log.Println("This is an informational message")

}

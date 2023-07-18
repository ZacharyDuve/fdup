package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var useTimestamp bool
	flag.BoolVar(&useTimestamp, "t", false, "Mark current timestamp as suffix and create duplicate")

	flag.Parse()

	fName := flag.Arg(0)

	if fName == "" {
		os.Stderr.WriteString("Error: Missing file name")
		printUsage()
		os.Exit(1)
	}

	fIn, err := os.Open(fName)

	if err != nil {
		os.Stderr.WriteString(fmt.Sprintln("Error occured opening source file", err))
		printUsage()
		os.Exit(2)
	}
	var fOut *os.File

	if useTimestamp {
		timeStamp := time.Now().UnixMilli()
		fOut, err = os.Create(fmt.Sprint(fName, ".", timeStamp))
	} else {
		fOut, err = os.Create(fName + ".orig")
	}

	if err != nil {
		os.Stderr.WriteString(fmt.Sprintln("Error occured opening new file", err))
		os.Exit(3)
	}

	_, err = io.Copy(fOut, fIn)

	if err != nil {
		os.Stderr.WriteString(fmt.Sprintln("Error occured copying file contents", err))
		os.Exit(3)
	}

	os.Exit(0)
}

func printUsage() {
	os.Stdout.WriteString("fdup [-t] filename\n")
	os.Stdout.WriteString("\t-t append timestamp of millis since epoch as suffix instead of \".orig\"\n")
}

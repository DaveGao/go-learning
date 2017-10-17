package main

import (
	"fmt"
	"os"
	"flag"
	"log"
	"syscall"
)

type selpgArgs struct {
	PStart int
	PEnd int
	SrcFilename string
	PLen int
	PType int
	PrintDest string
}
var progName string;

func usage() {
	fmt.Println("\nUSAGE:  %v -s StartPage -e EndPage [ -f | -l lines_per_page ] [ -d dest ] [ in_filename ]\n", progName);
}

func ErrorReminder(message string) {
	fmt.Println(progName,": ", message)
}

func processArgs(aArgs []string, selP *selpgArgs) {

	var ArgsLen = len(aArgs)
	var tSrc string
	var tPstart int
	flag.IntVar(&tPstart, "s", -1, "USAGE: -s StartPage")
	var tPEnd int
	flag.IntVar(&tPEnd, "e", -1, "USAGE: -e EndPage")
	var tPLen int
	flag.IntVar(&tPLen, "l", 72, "USAGE: -l lines_per_page")
	var tType = false
	flag.BoolVar(&tType, "f", true, "USAGE: -f variable_line")
	var tDest string
	flag.StringVar(&tDest, "d", "", "USAGE: -d Destination_output" )
	flag.Parse()

	var restArgs = flag.Args()
	if len(restArgs) == 1 {
		selP.SrcFilename = tSrc
	} else {
		ErrorReminder("unknown arguments")
		for i := 0; i < len(restArgs); i++ {
			fmt.Println(restArgs[i])
		}
	}

	if ArgsLen < 3 {
		ErrorReminder("not enough arguments")
		usage()
		os.Exit(1)
	}

	if tPstart < 1 {
		ErrorReminder("error StartPage number")
		usage()
		os.Exit(2)
	} else {
		selP.PStart = tPstart

	}

	if tPEnd < 1 || tPEnd < tPstart {
		ErrorReminder("error EndPage number")
		usage()
		os.Exit(3)
	} else {
		selP.PEnd = tPEnd
	}

	if tPLen < 1 {
		ErrorReminder("error PageLength number")
		usage()
		os.Exit(4)
	} else {
		selP.PLen = tPLen

	}
	if tType == true {
		selP.PType = 'f'

	} else {
		selP.PType = 'l'
	}
	selP.PrintDest = tDest
}

func processInput(selP *selpgArgs) {

	var fIn, fOut *os.File
	var currPage, currLine int

	if selP.SrcFilename == "" {
		fIn = os.Stdin
	} else {
		f, err := os.Open(selP.SrcFilename)
		defer f.Close()
		fIn = f
		if err != nil {
			log.Fatal(err)
			os.Exit(5)
		}
	}

	if selP.PrintDest == "" {
		fOut = os.Stdout
	} else {
		f, err := os.Open(selP.SrcFilename)
		defer f.Close()
		fOut = f
		if err != nil {
			log.Fatal(err)
			os.Exit(6)
		}
	}

	if selP.PType == 'l' {
		currPage = 1
		currLine = 0
	} else {
		currPage = 1
		for {
			if ()
		}
	}
	if currPage < selP.StartPage {

	} else if currPage < selP.EndPage {

	}
	if error(fIn) {
		ErrorReminder("")
	} else {
		ErrorReminder("done")
	}




}

func main() {
	tSelpgArgs := selpgArgs{-1, -1, "", 72, 'l' , ""}
	progName = os.Args[0]

	processArgs(os.Args[1:], &tSelpgArgs)
	processInput(&tSelpgArgs)

}

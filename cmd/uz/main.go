package main

import (
	"flag"

	"github.com/xyproto/textoutput"
	"github.com/xyproto/unzip"
)

const versionString = "uz 1.0.0"

func main() {
	o := textoutput.NewTextOutput(true, true)
	flag.Parse()
	if len(flag.Args()) == 0 {
		o.Printf("<blue>%s</blue>\n", versionString)
		o.Println("Provide a ZIP filename and (optionally) a directory to extract to.")
		return
	}
	zipfile := flag.Args()[0]
	directory := "."
	if len(flag.Args()) > 1 {
		directory = flag.Args()[1]
		o.Printf("<green>Extracting %s to %s...</green>\n", zipfile, directory)
	} else {
		o.Printf("<green>Extracting %s...</green>\n", zipfile)
	}
	if err := unzip.FilterExtract(zipfile, directory, func(filename string) bool {
		o.Printf("<blue>Extracting %s</blue>\n", filename)
		return true // continue
	}); err != nil {
		o.ErrExit(err.Error())
	}
	o.Println("<green>Done.</green>")
}

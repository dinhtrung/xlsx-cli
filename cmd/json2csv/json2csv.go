package main

import (
	"flag"
	"log"
	"os"
	"github.com/jscherff/json2csv/io"
)

const dstFile string =`out.csv`

var (
        fSrcUrl = flag.String("url", "", "Read JSON objects from `<url>`")
        fSrcFile = flag.String("i", "", "Read JSON objects from `<file>`")
        fDstFile = flag.String("o", dstFile, "Write CSV records to `<file>`")
)

func init() {

	log.SetFlags(0)
	flag.Parse()

	if *fSrcUrl == "" && *fSrcFile == "" {
		log.Print("You must specify a source file or URL")
		flag.Usage()
		os.Exit(1)
	}
}

func main() {

	var err error

	rw := io.NewReadWriter()

	switch true {
	case *fSrcUrl != "":
		err = rw.ReadUrl(*fSrcUrl)
	case *fSrcFile != "":
		err = rw.ReadFile(*fSrcFile)
	}

	if err == nil {
		err = rw.WriteFile(*fDstFile)
	}

	if err != nil {
		log.Fatal(err)
	}
}

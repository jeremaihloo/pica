package main

import (
	"github.com/alecthomas/kingpin"
	"os"
	"fmt"
	pica2 "github.com/jeremaihloo/pica"
)

var (
	app      = kingpin.New("pica", "A command line for api test and doc generate")
	filename = app.Arg("filename", "Api file.").ExistingFile()
	apiNames = app.Arg("apiNames", "Api names to excute").Strings()
	delay    = app.Flag("delay", "Delay after one api request.").Int()
	output   = app.Flag("output", "Output file.").OpenFile(os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	filetype = app.Flag("filetype", "The type of api file.").Default("pica").String()

	debug       = app.Flag("debug", "Debug mode.").Bool()
	ifRun       = app.Flag("run", "Run file.").Default("1").Bool()
	ifConvert   = app.Flag("convert", "Convert file.").Bool()
	ifDoc       = app.Flag("doc", "Generate document for a api file.").Bool()
	ifServer    = app.Flag("server", "Run as a document server").Bool()
	ifParse     = app.Flag("parse", "Parse api file.").Bool()
	ifFormat    = app.Flag("format", "Format api file.").Bool()
	docTemplate = app.Flag("template", "Document custom template.").ExistingFile()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	if !*debug {
		//defer func() {
		//	if err := recover(); err != nil {
		//		fmt.Printf("\nerror: %s\n", err)
		//	}
		//}()
	}
	if *filename == "" {
		kingpin.Usage()
	} else {
		fmt.Printf("\nRunning file:  %s\n\n", *filename)

		pica := pica2.NewPica(
			*filename,
			*output,
			*delay,
			*ifRun,
			*ifFormat,
			*ifConvert,
			*ifDoc,
			*ifServer,
			*docTemplate)
		err := pica.Run()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

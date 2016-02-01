package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
)

type Data struct {
	String string `json:"String" xml:"String"`
	Int    int    `json:"Int" xml:"Int"`
}

func main() {
	argsString := ""
	argsInt := 0
	// json
	jsonFormat := flag.NewFlagSet("json", flag.ExitOnError)
	jsonFormat.StringVar(&argsString, "string", "", "")
	jsonFormat.IntVar(&argsInt, "int", 0, "")
	// xml
	xmlFormat := flag.NewFlagSet("xml", flag.ExitOnError)
	xmlFormat.StringVar(&argsString, "string", "", "")
	xmlFormat.IntVar(&argsInt, "int", 0, "")
	if len(os.Args) < 2 {
		fmt.Println("Subcommands:")
		fmt.Println("")
		fmt.Println(" json")
		jsonFormat.PrintDefaults()
		fmt.Println(" xml")
		xmlFormat.PrintDefaults()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "json":
		jsonFormat.Parse(os.Args[2:])
	case "xml":
		xmlFormat.Parse(os.Args[2:])
	default:
		fmt.Println("Subcommands:")
		fmt.Println("")
		fmt.Println(" json")
		jsonFormat.PrintDefaults()
		fmt.Println(" xml")
		xmlFormat.PrintDefaults()
		os.Exit(1)
	}
	if jsonFormat.Parsed() {
		data := Data{argsString, argsInt}
		js, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Printf("%s\n", string(js))
	} else if xmlFormat.Parsed() {
		data := Data{argsString, argsInt}
		xl, err := xml.MarshalIndent(data, "", "  ")
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Printf("%s\n", string(xl))
	}
}

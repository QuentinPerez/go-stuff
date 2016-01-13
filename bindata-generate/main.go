package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/Sirupsen/logrus"
)

type GenObject struct {
	Object      string
	PackageName string
	Type        string
}

func main() {
	Name := flag.String("name", "", "Object Name")

	flag.Parse()
	if *Name == "" {
		flag.Usage()
		return
	}
	resource, err := Asset("templates/stringer.tmpl")
	if err != nil {
		logrus.Fatal(err)
	}
	tmpl, err := template.New("template").Parse(string(resource))
	if err != nil {
		logrus.Fatal(err)
	}
	fileName := fmt.Sprintf("%v.go", *Name)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		logrus.Fatal(err)
	}
	if err := tmpl.Execute(file, GenObject{Object: string((*Name)[0]), PackageName: "stringer", Type: *Name}); err != nil {
		logrus.Fatal(err)
	}
}

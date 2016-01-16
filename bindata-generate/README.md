### Bindata-generate

#### How to generate golang code with go generate and go-bindata


- We will create a generator of class which implements the `stringer` interface.

- First, we need to generate the `resource.go` with `go-bindata` and `stringer.tmpl`

```console
$> cat templates/stringer.tmpl
package {{ .PackageName }}

import "fmt"

func ({{ .Object }} {{ .Type }}) String() string {
    return "{{ .Type }}"
}
$> go-bindata -o resource.go templates/stringer.tmpl
```
- Then, the `main.go` file which generates a new file with the template

```golang
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
```

- Now, we will install the `bindata-generate` command

```console
$> go install
```

- To finish, we will use `go generate` with the test files

```console
$> cd test
$> cat generate.go
package generate

//go:generate bindata-generate -name hello
//go:generate bindata-generate -name world
$> go generate
```

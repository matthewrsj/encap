package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/matthewrsj/encap/animal"
)

// report is the template to output information to the command line
const report = `
weight: {{.Weight}} lbs
age:    {{.Age}} yrs

speak {{.Name}}! "{{.Speak}}"
eat {{.Name}}!   "{{.Eat}}"
bite {{.Name}}!  "{{.Bite}}"

{{.Name}} now weighs {{.Weight}} lbs
--------------------------------------------------

`

// Report takes an Animal interface and prints it out via the template report
func Report(a animal.Animal) {
	tp := strings.TrimPrefix(reflect.TypeOf(a).String(), "*animal.")
	fmt.Printf("%s is a %s that can speak, eat, and bark!\n", a.Name(), tp)
	t := template.Must(template.New("report").Parse(report))
	err := t.Execute(os.Stdout, a)
	if err != nil {
		log.Fatalf("error executing template: %v", err)
	}
}

func main() {
	// each animal is a different type, but all animals implement the animal.Animal interface
	animals := []animal.Animal{
		animal.NewDog("Rover", 50, 5),
		animal.NewShowDog("Princess", 5, 2),
		animal.NewChihuahua("Taco", 5, 1),
		animal.NewGopher("Henry", 3, 1),
	}

	// since each animal implements the animal.Animal interface they can all be passed to Report()
	for _, a := range animals {
		Report(a)
	}
}

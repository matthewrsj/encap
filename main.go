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

const report = `
weight: {{.Weight}} lbs
age:    {{.Age}} yrs

speak {{.Name}}! "{{.Speak}}"
eat {{.Name}}!   "{{.Eat}}"
bite {{.Name}}!  "{{.Bite}}"

{{.Name}} now weighs {{.Weight}} lbs
--------------------------------------------------

`

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
	animals := []animal.Animal{
		animal.NewDog("Rover", 50, 5),
		animal.NewShowDog("Princess", 5, 2),
		animal.NewChihuahua("Taco", 5, 1),
		animal.NewGopher("Henry", 3, 1),
	}

	for _, a := range animals {
		Report(a)
	}
}

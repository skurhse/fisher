//go:generate go run generate.go

package main

import (
	"log"
	"os"
	"text/template"
)

type Bound struct {
	Name     string
	IsLower  bool
	IsUpper  bool
	Index    int
	Polarity int
	Shift    int
}

type Coord struct {
	Name string
}

type Printer struct {
	Lower  Bound
	Upper  Bound
	Bounds [2]Bound
	Coords [2]Coord
}

var (
	coords = [2]Coord{
		{
			Name: "X",
		},
		{
			Name: "Y",
		},
	}

	lower = Bound{
		Name:     "Lower",
		IsLower:  true,
		IsUpper:  false,
		Index:    -1,
		Polarity: -1,
		Shift:    0,
	}

	upper = Bound{
		Name:     "Upper",
		IsLower:  false,
		IsUpper:  true,
		Index:    -1,
		Polarity: 1,
		Shift:    1,
	}

	printer = Printer{
		Lower:  lower,
		Upper:  upper,
		Bounds: [2]Bound{lower, upper},
		Coords: coords,
	}
)

func main() {
	template := template.Must(template.ParseFiles("../binsrch.tmpl"))

	out, err := os.Create("../binsrch.go")

	if err != nil {
		log.Fatalf("%v", err)
	}
	defer out.Close()

	err = template.Execute(out, printer)

	if err != nil {
		log.Fatalf("%v", err)
	}
}

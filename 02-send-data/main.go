package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func main() {
	tpl02 := `Hola {{ .Name }}, tienes {{ .Age }} años
`
	t, err := template.New("ejercicio02").Parse(tpl02)
	if err != nil {
		log.Fatalf("Error al hacer parse del template %v", err)
	}

	p := person{"EDteam", 5}
	err = t.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalf("Error al ejecutar el template %v", err)
	}
}

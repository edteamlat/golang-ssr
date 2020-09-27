package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl04 := `Hola
	{{- . -}}
	cómo estás?
`
	t, err := template.New("ejercicio04").Parse(tpl04)
	if err != nil {
		log.Fatalf("Error al hacer parse del template %v", err)
	}

	err = t.Execute(os.Stdout, "EDteam")
	if err != nil {
		log.Fatalf("Error al ejecutar el template %v", err)
	}
}

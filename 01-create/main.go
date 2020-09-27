package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl01 := "Hola mundo"
	t, err := template.New("ejercicio01").Parse(tpl01)
	if err != nil {
		log.Fatalf("Error al hacer parse del template %v", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("Error al ejecutar el template %v", err)
	}
}

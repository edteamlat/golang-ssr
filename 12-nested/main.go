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
	t, err := template.New("").ParseGlob("templates/*.tpl")
	if err != nil {
		log.Fatalf("Error al hacer parse del template %v", err)
	}

	p := person{"María", 20}
	err = t.ExecuteTemplate(os.Stdout, "wrapper", &p)
	if err != nil {
		log.Fatalf("Error al ejecutar el template invitación %v", err)
	}
}

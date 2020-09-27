package main

import (
	"fmt"
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
	err = t.ExecuteTemplate(os.Stdout, "invitation.tpl", &p)
	if err != nil {
		log.Fatalf("Error al ejecutar el template invitación %v", err)
	}

	if p.Age >= 18 {
		fmt.Println("************************")
		err = t.ExecuteTemplate(os.Stdout, "confirmation.tpl", &p)
		if err != nil {
			log.Fatalf("Error al ejecutar el template confirmación %v", err)
		}
	}
}

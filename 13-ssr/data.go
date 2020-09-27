package main

type class struct {
	Title    string
	Duration string
}

type module struct {
	Title       string
	Description string
	Classes     []class
}

type course struct {
	Slug         string
	Img          string
	Title        string
	Name         string
	Description  string
	Average      float64
	Professor    string
	ProfessorImg string
	Price        float64
	Modules      []module
}

type gridPage struct {
	InternalTemplate string
	Courses          []course
}

type coursePage struct {
	InternalTemplate string
	Course           course
}

func loadGrid() []course {
	return []course{
		{
			Slug:         "go",
			Img:          "https://edteam-media.s3.amazonaws.com/courses/big/91e149d0-961a-4594-a8ff-0a625be9cdd2.png",
			Title:        "Go desde cero",
			Name:         "Go desde cero",
			Description:  "Adquiere los conocimientos básicos para iniciar a programar con Go.",
			Average:      4.7,
			Professor:    "Alejo Rodriguez",
			ProfessorImg: "https://edteam-media.s3.amazonaws.com/users/thumbnail/952327c3-2bd9-41d1-819e-9b5d7eb84c13.jpg",
			Price:        30,
			Modules: []module{
				{
					Title:       "Qué es Go",
					Description: "Aprenderemos qué es el lenguaje Go",
					Classes: []class{
						{
							Title:    "Historia",
							Duration: "05:33",
						},
						{
							Title:    "Creadores",
							Duration: "03:02",
						},
					},
				},
				{
					Title:       "Sintaxis",
					Description: "Bases del lenguaje",
					Classes: []class{
						{
							Title:    "Declaración de Variables",
							Duration: "05:33",
						},
						{
							Title:    "Constantes",
							Duration: "03:02",
						},
					},
				},
			},
		},
		{
			Slug:         "go-poo",
			Img:          "https://edteam-media.s3.amazonaws.com/courses/big/8f37ebcc-16a1-4001-93e1-5f21893cd3cc.jpg",
			Title:        "POO con Go",
			Name:         "POO con Go",
			Description:  "Aprende a usar la programación orientada a objetos en Go",
			Average:      4.8,
			Professor:    "Alejo Rodriguez",
			ProfessorImg: "https://edteam-media.s3.amazonaws.com/users/thumbnail/952327c3-2bd9-41d1-819e-9b5d7eb84c13.jpg",
			Price:        30,
		},
		{
			Slug:         "go-database",
			Img:          "https://edteam-media.s3.amazonaws.com/courses/big/4d60ef81-2e58-457f-97c7-ee8847663985.jpg",
			Title:        "Bases de datos con Go",
			Name:         "Bases de datos con Go",
			Description:  "Aprende a integrar y usar las bases de datos en Go",
			Average:      4.8,
			Professor:    "Alejo Rodriguez",
			ProfessorImg: "https://edteam-media.s3.amazonaws.com/users/thumbnail/952327c3-2bd9-41d1-819e-9b5d7eb84c13.jpg",
			Price:        30,
		},
		{
			Slug:         "go-testing",
			Img:          "https://edteam-media.s3.amazonaws.com/courses/big/a9913502-8af2-400b-8095-7b78f52200dc.png",
			Title:        "Testing con Go",
			Name:         "Testing con Go",
			Description:  "Aprende a crear tests y medir la eficiencia de tus proyectos con Go",
			Average:      4.7,
			Professor:    "Alexys Lozada",
			ProfessorImg: "https://edteam-media.s3.amazonaws.com/users/thumbnail/66f015b2-0dfb-4ba9-bd0d-f7a7e1855275.jpeg",
			Price:        24,
		},
	}
}

func getCourse(slug string) course {
	data := loadGrid()
	for _, v := range data {
		if v.Slug == slug {
			return v
		}
	}

	return course{}
}

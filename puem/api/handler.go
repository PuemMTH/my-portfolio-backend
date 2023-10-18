package api

import (
	"time"

	"puem/puem/types"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutesAPI(app *fiber.App) {
	app.Get("/api/puem", func(c *fiber.Ctx) error {
		time.Sleep(1 * time.Second)

		return c.JSON(types.Data{
			Profile: types.ProfileData{
				Name:            "Tanapat Eiam-arj",
				Nickname:        "puem",
				Phone:           "061-654-4123",
				Email:           "tanapat.e@ku.th",
				Github:          "github.com/PuemMTH",
				Linkedin:        "linkedin.com/in/puem-methawee",
				FullVersionLink: "https://puemmth.github.io/my-portfolio",
				Location:        "Bangkok, Thailand",
				Website:         "https://puemmth.github.io/my-portfolio",
			},
			Skills: []types.Skill{
				{
					Name:        "Backend",
					Description: "Node.js, Next.js, Lavavel, golang, Express.js, MongoDB, MySQL, Redis, Docker",
				},
				{
					Name:        "Frontend",
					Description: "React.js, HTML, CSS, JavaScript, TypeScript, TailwindCSS",
				},
				{
					Name:        "DevOps",
					Description: "AWS, GCP, Azure, CI/CD, Terraform, Jenkins",
				},
				{
					Name:        "Mobile",
					Description: "React Native, Flutter",
				},
				{
					Name:        "Other",
					Description: "Git, Linux, Nginx, Apache, Confluence, Agile, Scrum, Kanban, Trello",
				},
			},
			Projects: []types.Project{
				{
					Name:         "Portfolio Website",
					Description:  "This website portfolio is made by Next.js and TailwindCSS",
					Link:         "https://puem.tech/",
					Technologies: []string{"Next.js", "TailwindCSS"},
					Images:       []string{"https://puem.tech/"},
					Date:         "2023",
				},
				{
					Name:         "Table Website",
					Description:  "This website is table website for manage student in class",
					Link:         "https://table.puem.tech/",
					Technologies: []string{"Next.js", "TailwindCSS", "Node.js", "MongoDB"},
					Images:       []string{"https://puem.tech/"},
					Date:         "2022",
				},
				{
					Name:         "Manage Credit Website",
					Description:  "This website is manage credit website for manage credit in class",
					Link:         "https://credit.puem.tech/",
					Technologies: []string{"Next.js", "TailwindCSS", "Node.js", "JSON"},
					Images:       []string{"https://puem.tech/"},
					Date:         "2021",
				},
			},
			Experiences: []types.Experience{
				{
					Name:        "KUNLUN LAB",
					Description: "Software Engineer Intern",
				},
			},
			Educations: []types.EducationData{
				{
					Name:        "Kasetsart University",
					Description: "Computer Engineering",
					Date:        "2021 - Present",
				},
			},
			Certifications: []types.Certificate{
				{
					Name:        "Matlab Onramp",
					Description: "Mathworks",
					Date:        "2021",
					Link:        "https://matlabacademy.mathworks.com/",
				},
			},
			Languages: []string{
				"JavaScript",
				"TypeScript",
				"HTML",
				"CSS",
				"PHP",
				"Python",
				"Java",
				"C",
				"Go",
				"Dart",
			},
			Interests: []string{
				"Technology",
				"Coding",
				"Music",
			},
		})
	})
}

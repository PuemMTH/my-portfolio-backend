package api

import (
	"puem/puem/databases/models"

	"puem/puem/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutesDatabase(app *fiber.App, db *gorm.DB) {

	app.Get("api/profile", middleware.Auth, func(c *fiber.Ctx) error {
		// app.Get("api/profile", func(c *fiber.Ctx) error {
		var profile []models.ProfileData
		if err := db.Preload("Github").Preload("Experience").Preload("Certificate").Preload("Educations").Preload("Projects").Preload("Skills").Preload("Website").Preload("Languages").Preload("Interests").Find(&profile).Error; err != nil {
			return c.JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(profile)
	})

	app.Get("api/profile/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")              // คือ การรับค่า id ที่ส่งมาจาก url
		var profiles []models.ProfileData // คือ ตัวแปรที่เก็บข้อมูลที่ได้จากฐานข้อมูล

		if err := db.Preload("Github").Preload("Experience").Preload("Certificate").Preload("Educations").Preload("Projects").Preload("Skills").Preload("Website").Preload("Languages").Preload("Interests").Where("id = ?", id).Find(&profiles).Error; err != nil {
			return c.JSON(fiber.Map{"error": err.Error()})
		}

		var Github []string
		for _, github := range profiles[0].Github {
			Github = append(Github, github.Github)
		}

		return c.JSON(Github)

	})

	app.Get("api/add", func(c *fiber.Ctx) error {
		profile := models.ProfileData{
			Name:     "Tanapat Eiam-arj",
			Nickname: "puem",
			Phone:    "061-654-4123",
			Email:    "tanapat.e@ku.th",
			Linkedin: "linkedin.com/in/puem-methawee",
			Location: "Bangkok, Thailand",
			Github: []models.Github{
				{Github: "github.com/PuemMTH"},
				{Github: "github.com/puemku"},
			},
			Experience: []models.Experience{
				{
					Name:        "KUNLUN LAB",
					Description: "Software Engineer Intern",
				},
			},
			Certificate: []models.Certificate{
				{
					Name:        "Matlab Onramp",
					Description: "Mathworks",
					Link:        "https://matlabacademy.mathworks.com/",
					Date:        "2021",
				},
			},
			Educations: []models.EducationData{
				{
					NameEdu:     "Kasetsart University",
					Description: "Computer Engineering",
					Date:        "2021 - Present",
				},
			},
			Projects: []models.Projects{
				{
					Name:        "Portfolio Website",
					Description: "This website portfolio is made by Next.js and TailwindCSS",
					Link:        "https://puem.tech/",
					Technologies: []models.Technologies{
						{Technologies: "Next.js"},
						{Technologies: "TailwindCSS"},
					},
					Images: []models.Images{
						{Images: "https://puem.tech/"},
					},
					Date: "2023",
				},
			},
			Skills: []models.Skills{
				{
					Name:        "Backend",
					Description: "Node.js, Next.js, Lavavel, golang, Express.js, MongoDB, MySQL, Redis, Docker",
				},
			},
			Website: []models.Website{
				{Website: "https://puem.tech/"},
				{Website: "https://ku.puem.tech/"},
				{Website: "https://ask.puem.tech/"},
			},
			Languages: []models.Languages{
				{Languages: "JavaScript"},
				{Languages: "TypeScript"},
			},
			Interests: []models.Interests{
				{Interests: "Technology"},
				{Interests: "Coding"},
			},
		}
		db.Create(&profile)
		return c.JSON(profile)
	})
}

// return c.JSON(types.Data{
// 	Profile: types.ProfileData{
// 		Name:            "Tanapat Eiam-arj",
// 		Nickname:        "puem",
// 		Phone:           "061-654-4123",
// 		Email:           "tanapat.e@ku.th",
// 		Github:          "github.com/PuemMTH",
// 		Linkedin:        "linkedin.com/in/puem-methawee",
// 		FullVersionLink: "https://puemmth.github.io/my-portfolio",
// 		Location:        "Bangkok, Thailand",
// 		Website:         "https://puemmth.github.io/my-portfolio",
// 	},
// 	Skills: []types.Skill{
// 		{
// 			Name:        "Backend",
// 			Description: "Node.js, Next.js, Lavavel, golang, Express.js, MongoDB, MySQL, Redis, Docker",
// 		},
// 		{
// 			Name:        "Frontend",
// 			Description: "React.js, HTML, CSS, JavaScript, TypeScript, TailwindCSS",
// 		},
// 		{
// 			Name:        "DevOps",
// 			Description: "AWS, GCP, Azure, CI/CD, Terraform, Jenkins",
// 		},
// 		{
// 			Name:        "Mobile",
// 			Description: "React Native, Flutter",
// 		},
// 		{
// 			Name:        "Other",
// 			Description: "Git, Linux, Nginx, Apache, Confluence, Agile, Scrum, Kanban, Trello",
// 		},
// 	},
// 	Projects: []types.Project{
// 		{
// 			Name:         "Portfolio Website",
// 			Description:  "This website portfolio is made by Next.js and TailwindCSS",
// 			Link:         "https://puem.tech/",
// 			Technologies: []string{"Next.js", "TailwindCSS"},
// 			Images:       []string{"https://puem.tech/"},
// 			Date:         "2023",
// 		},
// 		{
// 			Name:         "Table Website",
// 			Description:  "This website is table website for manage student in class",
// 			Link:         "https://table.puem.tech/",
// 			Technologies: []string{"Next.js", "TailwindCSS", "Node.js", "MongoDB"},
// 			Images:       []string{"https://puem.tech/"},
// 			Date:         "2022",
// 		},
// 		{
// 			Name:         "Manage Credit Website",
// 			Description:  "This website is manage credit website for manage credit in class",
// 			Link:         "https://credit.puem.tech/",
// 			Technologies: []string{"Next.js", "TailwindCSS", "Node.js", "JSON"},
// 			Images:       []string{"https://puem.tech/"},
// 			Date:         "2021",
// 		},
// 	},
// 	Experiences: []types.Experience{
// 		{
// 			Name:        "KUNLUN LAB",
// 			Description: "Software Engineer Intern",
// 		},
// 	},
// 	Educations: []types.EducationData{
// 		{
// 			Name:        "Kasetsart University",
// 			Description: "Computer Engineering",
// 			Date:        "2021 - Present",
// 		},
// 	},
// 	Certifications: []types.Certificate{
// 		{
// 			Name:        "Matlab Onramp",
// 			Description: "Mathworks",
// 			Date:        "2021",
// 			Link:        "https://matlabacademy.mathworks.com/",
// 		},
// 	},
// 	Languages: []string{
// 		"JavaScript",
// 		"TypeScript",
// 		"HTML",
// 		"CSS",
// 		"PHP",
// 		"Python",
// 		"Java",
// 		"C",
// 		"Go",
// 		"Dart",
// 	},
// 	Interests: []string{
// 		"Technology",
// 		"Coding",
// 		"Music",
// 	},
// })

package types

type ProfileData struct {
	Name            string `json:"name"`
	Nickname        string `json:"nickname"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	Github          string `json:"github"`
	Linkedin        string `json:"linkedin"`
	FullVersionLink string `json:"fullVersionLink"`
	Location        string `json:"location"`
	Website         string `json:"website"`
}

type EducationData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

type Certificate struct {
	Name        string `json:"name"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type Experience struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Project struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Link         string   `json:"link"`
	Technologies []string `json:"technologies"`
	Images       []string `json:"images"`
	Date         string   `json:"date"`
}

type Skill struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Data struct {
	Profile        ProfileData     `json:"profile"`
	Skills         []Skill         `json:"skills"`
	Projects       []Project       `json:"projects"`
	Experiences    []Experience    `json:"experiences"`
	Educations     []EducationData `json:"educations"`
	Certifications []Certificate   `json:"certifications"`
	Languages      []string        `json:"languages"`
	Interests      []string        `json:"interests"`
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

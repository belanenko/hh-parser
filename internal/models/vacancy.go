package models

type Vacancy struct {
	Description        string             `json:"description"`
	DatePosted         string             `json:"datePosted"`
	Title              string             `json:"title"`
	HiringOrganization HiringOrganization `json:"hiringOrganization"`
	ValidThrough       string             `json:"validThrough"`
	JobLocation        JobLocation        `json:"jobLocation"`
	EmploymentType     string             `json:"employmentType"`
	Industry           []string           `json:"industry"`
	BaseSalary         BaseSalary         `json:"baseSalary"`
	Identifier         Identifier         `json:"identifier"`
}

type BaseSalary struct {
	Currency string `json:"currency"`
	Value    Value  `json:"value"`
}

type Value struct {
	UnitText string `json:"unitText"`
	MinValue int64  `json:"minValue"`
	MaxValue int64  `json:"manValue"`
}

type HiringOrganization struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type Identifier struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type JobLocation struct {
	Address Address `json:"address"`
}

type Address struct {
	AddressLocality string `json:"addressLocality"`
	AddressRegion   string `json:"addressRegion"`
	AddressCountry  string `json:"addressCountry"`
}

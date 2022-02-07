package models

import "time"

type Vacancy struct {
	Context            string    `json:"@context"`
	Type               string    `json:"@type"`
	Description        string    `json:"description"`
	DatePosted         time.Time `json:"datePosted"`
	Title              string    `json:"title"`
	HiringOrganization struct {
		Type string `json:"@type"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"hiringOrganization"`
	ValidThrough time.Time `json:"validThrough"`
	JobLocation  struct {
		Type    string `json:"@type"`
		Address struct {
			Type            string `json:"@type"`
			AddressLocality string `json:"addressLocality"`
			AddressRegion   string `json:"addressRegion"`
			AddressCountry  string `json:"addressCountry"`
			StreetAddress   string `json:"streetAddress"`
		} `json:"address"`
	} `json:"jobLocation"`
	JobLocationType string   `json:"jobLocationType"`
	EmploymentType  string   `json:"employmentType"`
	Industry        []string `json:"industry"`
	BaseSalary      struct {
		Type     string `json:"@type"`
		Currency string `json:"currency"`
		Value    struct {
			Type     string `json:"@type"`
			UnitText string `json:"unitText"`
			MinValue int    `json:"minValue"`
			MaxValue int    `json:"maxValue"`
		} `json:"value"`
	} `json:"baseSalary"`
	Identifier struct {
		Type  string `json:"@type"`
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"identifier"`
}

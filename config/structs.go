package config

type University struct {
	Name        string   `json:"name"`
	CountryName string   `json:"country"`
	IsoCode     string   `json:"alpha_two_code"`
	WebPages    []string `json:"web_pages"`
	Map         struct {
		OpenStreetMaps string `json:"openstreetmaps"`
	} `json:"maps"`
	Languages map[string]string `json:"languages"`
}

type Country struct {
	CountryName struct {
		Name string `json:"common"`
	} `json:"name"`
	IsoCode    string   `json:"cca2"`
	Neighbours []string `json:"borders"`
	Map        struct {
		OpenStreetMaps string `json:"openstreetmaps"`
	} `json:"maps"`
	Languages map[string]string `json:"languages"`
}

type Diagnostics struct {
	UniversityStatus int    `json:"universitiesapi"`
	CountryStatus    int    `json:"countriesapi"`
	Version          string `json:"version"`
	UpTime           int    `json:"uptime"`
}

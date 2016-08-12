package main

type Config struct {
	Client    Client    `json:"client"`
	Dashboard Dashboard `json:"dashboard"`
}

type Client struct {
	ApiKey   string `json:"apiKey"`
	BaseURL  string `json:"baseUrl"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Dashboard struct {
	Title          string `json:"title"`
	Interval       string `json:"interval"`
	Timezone       string `json:"timezone"`
	TemplateType   string `json:"templateType"` //web or local
	InheritTitle   string `json:"inheritTitle"`
	InheritPanelID int64  `json:"inheritPanelid"`
	TemplateFile   string `json:"templateFile"`
	Rows           []Row  `json:"rows"`
}

type Row struct {
	Title  string  `json:"title"`
	Panels []Panel `json:"panels"`
}

type Panel struct {
	Title  string   `json:"title"`
	DsType string   `json:"dsType"`
	Sqls   []string `json:"sqls"`
}

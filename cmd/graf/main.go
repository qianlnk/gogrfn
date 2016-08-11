package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/qianlnk/gogrfn"
)

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
	Timezone       string `json:"timezone"`
	InheritTitle   string `json:"inheritTitle"`
	InheritPanelID int64  `json:"inheritPanelid"`
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

func getConfig(filename string) (Config, error) {
	var config Config
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func main() {
	file := flag.String("file", "default.json", "Use -file <template json file>")
	flag.Parse()
	cfg, err := getConfig(*file)
	if err != nil {
		fmt.Println(err)
		return
	}
	cli, err := gografana.NewGrafanaClient(cfg.Client.ApiKey, cfg.Client.BaseURL, cfg.Client.User, cfg.Client.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	//get inherit panel
	inheritPanel, err := cli.GetInheritPanel(cfg.Dashboard.InheritTitle, cfg.Dashboard.InheritPanelID)
	if err != nil {
		fmt.Println(err)
		return
	}

	//get db, create db if it not exist
	var db *gografana.Dashboard
	db, err = cli.GetDashboard(cfg.Dashboard.Title)
	if err != nil {
		db, err = cli.NewDashboard(cfg.Dashboard.Title, cfg.Dashboard.Timezone)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for _, r := range cfg.Dashboard.Rows {
		//add new row
		db.AddRow(r.Title)
		for _, p := range r.Panels {
			//add new panel
			panelID := db.GetNextPanelID()
			db.Rows[len(db.Rows)-1].AddPanel(inheritPanel, panelID, p.Title, p.Sqls, p.DsType)
		}
	}

	//upload db
	err = cli.UploadDashboard(db)
	if err != nil {
		fmt.Println(err)
		return
	}

}

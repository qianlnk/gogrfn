package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/qianlnk/graf"
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
	Template       string `json:"template"`
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

func getClientCfg() (Client, error) {
	var cli Client
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return cli, err
	}

	err = json.Unmarshal(file, &cli)
	if err != nil {
		return cli, err
	}

	return cli, nil
}

func getDashboard(filename string) (Dashboard, error) {
	var db Dashboard
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return Dashboard{}, err
	}

	err = json.Unmarshal(file, &db)
	if err != nil {
		return Dashboard{}, err
	}

	return db, nil
}

func getTemplate(filename string) {}
func main() {
	file := flag.String("file", "default.json", "Use -file <template json file>")
	flag.Parse()
	cliCfg, err := getClientCfg()
	if err != nil {
		fmt.Println(err)
		return
	}
	cli, err := gografana.NewGrafanaClient(cliCfg.ApiKey, cliCfg.BaseURL, cliCfg.User, cliCfg.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	dbCfg, err := getDashboard(*file)
	if err != nil {
		fmt.Println(err)
		return
	}
	//get inherit panel
	inheritPanel, err := cli.GetInheritPanel(dbCfg.InheritTitle, dbCfg.InheritPanelID)
	if err != nil {
		fmt.Println(err)
		return
	}

	//get db, create db if it not exist
	var db *gografana.Dashboard
	db, err = cli.GetDashboard(dbCfg.Title)
	if err != nil {
		db, err = cli.NewDashboard(dbCfg.Title, dbCfg.Timezone)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for _, r := range dbCfg.Rows {
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

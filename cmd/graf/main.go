package main

import (
	"fmt"

	graf "github.com/qianlnk/graf"
	"github.com/voxelbrain/goptions"
)

func main() {
	options := struct {
		File string        `goptions:"-f, --file, description='The json file to add graph'"`
		Help goptions.Help `goptions:"-h, --help, description='Show this help'"`

		goptions.Verbs
		Export struct {
			Title   string `goptions:"-t, --title, obligatory, description='db title to export'"`
			PanelID int64  `goptions:"-p, --panel, description='panel id to export'"`
			Name    string `goptions:"-n, --name, obligatory, description='the template name to save'"`
		} `goptions:"export"`
	}{ // Default values goes here
	}
	goptions.ParseAndFail(&options)

	cliCfg, err := getClientCfg()
	if err != nil {
		fmt.Println(err)
		return
	}
	cli, err := graf.NewGrafanaClient(cliCfg.ApiKey, cliCfg.BaseURL, cliCfg.User, cliCfg.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	if options.File != "" {
		fmt.Println(newGraph(cli, options.File))
	} else if options.Verbs == "" {
		fmt.Println("Not enough param, pelase run 'graf －h' for help")
	} else {
		if options.Export.PanelID == 0 {
			fmt.Println(exportDashboard(cli, options.Export.Title, options.Export.Name))
		} else {
			fmt.Println(exportTemplate(cli, options.Export.Title, options.Export.PanelID, options.Export.Name))
		}
	}

}

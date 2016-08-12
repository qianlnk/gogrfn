package main

import (
	"errors"

	graf "github.com/qianlnk/graf"
)

func newGraph(cli *graf.Client, dbfile string) error {
	dbCfg, err := getDashboard(dbfile)
	if err != nil {
		return err
	}
	//get inherit panel
	var inheritPanel *graf.Panel
	if dbCfg.TemplateType == "web" {
		inheritPanel, err = cli.GetInheritPanel(dbCfg.InheritTitle, dbCfg.InheritPanelID)
		if err != nil {
			return err
		}
	} else {
		inheritPanel, err = getTemplate(dbCfg.TemplateFile)
		if err != nil {
			return err
		}
	}

	//get db, create db if it not exist
	var db *graf.Dashboard
	db, err = cli.GetDashboard(dbCfg.Title)
	if err != nil {
		db, err = cli.NewDashboard(dbCfg.Title, dbCfg.Interval, dbCfg.Timezone)
		if err != nil {
			return err
		}
	}

	for _, r := range dbCfg.Rows {
		//add new row
		addNewRow := true
		for _, p := range r.Panels {
			deleteRepetition(db, p.Title)

			if addNewRow {
				db.AddRow(r.Title)
				addNewRow = false
			}
			//add new panel
			panelID := db.GetNextPanelID()
			db.Rows[len(db.Rows)-1].AddPanel(inheritPanel, panelID, p.Title, p.Sqls, p.DsType)
		}
	}

	//upload db
	err = cli.UploadDashboard(db)
	if err != nil {
		return err
	}
	return errors.New("ok")
}

func deleteRepetition(db *graf.Dashboard, title string) {
	repRow := -1
	repPanel := -1
	for i, r := range db.Rows {
		for j, p := range r.Panels {
			if p.Title == title {
				repRow = i
				repPanel = j
			}
		}
	}

	if repRow != -1 {
		db.Rows[repRow].Panels = append(db.Rows[repRow].Panels[:repPanel], db.Rows[repRow].Panels[repPanel+1:]...)
		if len(db.Rows[repRow].Panels) == 0 {
			db.Rows = append(db.Rows[:repRow], db.Rows[repRow+1:]...)
		}
	}
}

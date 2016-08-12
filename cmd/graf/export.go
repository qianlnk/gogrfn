package main

import (
	"encoding/json"
	"errors"
	"fmt"

	graf "github.com/qianlnk/graf"
)

func exportTemplate(cli *graf.Client, title string, pid int64, name string) error {
	panel, err := cli.GetInheritPanel(title, pid)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(panel, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	filename := fmt.Sprintf("templates/%s", name)
	err = save(data, filename)
	if err != nil {
		return err
	}
	return errors.New("ok")
}

func exportDashboard(cli *graf.Client, title string, name string) error {
	db, err := cli.GetDashboard(title)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(db, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	filename := fmt.Sprintf("dashboards/%s", name)
	err = save(data, filename)
	if err != nil {
		return err
	}
	return errors.New("ok")
}

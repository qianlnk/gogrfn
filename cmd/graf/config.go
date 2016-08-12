package main

import (
	"encoding/json"
	"io/ioutil"

	graf "github.com/qianlnk/graf"
)

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

func getTemplate(filename string) (*graf.Panel, error) {
	panel := new(graf.Panel)

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return panel, err
	}

	err = json.Unmarshal(file, panel)
	if err != nil {
		return panel, err
	}

	return panel, nil
}

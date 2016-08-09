package main

import (
	"fmt"
	"testing"
)

const (
	key      = "eyJrIjoiZjRjTmZQYmZSVjg5YWJVV0I1anc0b0d0SlVBREtzUlAiLCJuIjoibG5rIiwiaWQiOjF9"
	URL      = "http://192.168.99.100:3000"
	user     = "admin"
	password = "admin"
)

//can't get orgs due to permission denied
func TestGetOrgs(t *testing.T) {
	cli, _ := NewGrafanaClient(key, URL, user, password)
	fmt.Println(cli.Orgs())
}

func TestNewOrg(t *testing.T) {
	cli, _ := NewGrafanaClient(key, URL, user, password)
	fmt.Println(cli.NewOrg("lnktest"))
}

func TestNewDashBoard(t *testing.T) {
	dbjson := `{"dashboard": {"id": null,"title": "Production erer","tags": [ "test" ],"timezone": "browser","rows": [{}],"schemaVersion": 6,"version": 0},"overwrite": false}`
	cli, _ := NewGrafanaClient(key, URL, user, password)
	fmt.Println(cli.NewDashBoard(dbjson))
}

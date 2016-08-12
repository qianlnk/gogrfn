package gografana

import (
	"encoding/json"
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

func TestGetDashboard(t *testing.T) {
	cli, _ := NewGrafanaClient(key, URL, user, password)
	fmt.Println(cli.GetDashboard("market-by-minute"))
}

func TestNewDashboard(t *testing.T) {
	cli, _ := NewGrafanaClient(key, URL, user, password)
	fmt.Println(cli.NewDashboard("lnk111", "6h", "utc"))
}

func TestDeleteDashboard(t *testing.T) {
	cli, _ := NewGrafanaClient(key, URL, user, password)
	fmt.Println(cli.DeleteDashboard(" "))
}

func TestAddNewRow(t *testing.T) {
	//Authorization
	cli, _ := NewGrafanaClient(key, URL, user, password)
	//get the dashboard you want upload
	db, err := cli.GetDashboard("market-by-minute")
	if err != nil {
		fmt.Println(err)
		return
	}
	//get the inherit panel
	panel, err := cli.GetInheritPanel("market-by-minute", 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	//add a new row
	db.AddRow("myrow1")

	//get next panelID
	panelID := db.GetNextPanelID()
	sqls := []string{`SELECT sum("Kafka") FROM "m.request" WHERE $timeFilter GROUP BY time(1h) fill(null)`,
		`SELECT sum("Send") FROM "m.request" WHERE $timeFilter GROUP BY time(1h) fill(null)`,
		`SELECT sum("Mysql") FROM "m.request" WHERE $timeFilter GROUP BY time(1h) fill(null)`}
	db.Rows[len(db.Rows)-1].AddPanel(panel, panelID, "m.request by hour", sqls, "influxdb")
	body, _ := json.Marshal(db)
	fmt.Println("add panel", string(body))
	cli.UploadDashboard(db)
}

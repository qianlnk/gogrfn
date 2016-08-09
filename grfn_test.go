package gografana

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
	dbjson := `{
    "dashboard": {
        "annotations": {
            "list": [ ]
        }, 
        "editable": true, 
        "hideControls": false, 
        "id": null, 
        "links": [ ], 
        "originalTitle": "Market By Minute", 
        "refresh": "30s", 
        "rows": [
            {
                "collapse": false, 
                "editable": true, 
                "height": "250px", 
                "panels": [ ], 
                "title": "New row"
            }
        ], 
        "schemaVersion": 12, 
        "sharedCrosshair": false, 
        "style": "dark", 
        "tags": [ ], 
        "templating": {
            "list": [ ]
        }, 
        "time": {
            "from": "now-3h", 
            "to": "now"
        }, 
        "timepicker": {
            "refresh_intervals": [
                "5s", 
                "10s", 
                "30s", 
                "1m", 
                "5m", 
                "15m", 
                "30m", 
                "1h", 
                "2h", 
                "1d"
            ], 
            "time_options": [
                "5m", 
                "15m", 
                "1h", 
                "6h", 
                "12h", 
                "24h", 
                "2d", 
                "7d", 
                "30d"
            ]
        }, 
        "timezone": "utc", 
        "title": "Market By Minute", 
        "version": 0
    }
}`
	cli, _ := NewGrafanaClient(key, URL, user, password)
	fmt.Println(cli.NewDashBoard(dbjson))
}

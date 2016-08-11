package gografana

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

func NewDashboard(title string, timezone string) *Dashboard {
	return &Dashboard{
		Editable:        true,
		HideControls:    false,
		OriginalTitle:   title,
		Refresh:         "30s",
		SchemaVersion:   12,
		SharedCrosshair: false,
		Time:            NewTime(),
		Timepicker:      NewTimepicker(),
		Timezone:        timezone,
		Title:           title,
		Version:         0,
	}
}

func NewTime() Time {
	return Time{
		From: "now-3h",
		To:   "now",
	}
}

func NewTimepicker() Timepicker {
	return Timepicker{
		Refresh_intervals: []string{"5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"},
		Time_options:      []string{"5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d"},
	}
}

func (d *Dashboard) AddRow(title string) {
	row := &Row{
		Collapse: false,
		Editable: true,
		Height:   "250px",
		Title:    title,
	}
	d.Rows = append(d.Rows, row)
}

func NewTarget(sql string, dsType string) Target {
	return Target{
		DsType:       dsType,
		Policy:       "default",
		Query:        sql,
		RawQuery:     true,
		ResultFormat: "time_series",
	}

}

func (d *Dashboard) GetNextPanelID() int64 {
	var maxID int64
	for _, r := range d.Rows {
		for _, p := range r.Panels {
			if maxID < p.ID {
				maxID++
			}
		}
	}
	return maxID + 1
}

func (c *Client) GetInheritPanel(title string, panelID int64) (*Panel, error) {
	fromDB, err := c.GetDashboard(title)
	if err != nil {
		return nil, err
	}

	panel := new(Panel)
	exist := false
	for _, r := range fromDB.Rows {
		for _, p := range r.Panels {
			if p.ID == panelID {
				*panel = *p
				exist = true
			}
		}
	}
	if exist == false {
		return panel, errors.New("Panel not found")
	}

	//clean this panel's Data
	panel.ID = -1
	panel.Targets = nil
	return panel, nil
}

func (p *Panel) GetNextRefId() string {
	if len(p.Targets) <= 0 {
		return "A"
	}
	var maxID byte
	for _, t := range p.Targets {
		tmpID := []byte(t.RefId)[0]
		if maxID < tmpID {
			maxID = tmpID
		}
	}
	maxID++

	return string(maxID)
}

func (r *Row) AddPanel(panel *Panel, panelID int64, title string, sqls []string, dsType string) error {
	pnl := new(Panel)
	*pnl = *panel
	pnl.ID = panelID
	pnl.Title = title
	for _, sql := range sqls {
		target := NewTarget(sql, dsType)
		target.RefId = pnl.GetNextRefId()
		pnl.Targets = append(pnl.Targets, target)
	}

	r.Panels = append(r.Panels, pnl)
	return nil
}

func (c *Client) NewDashboard(title, timezone string) (*Dashboard, error) {
	var dbul DashboardUploader
	dbul.Dashboard = NewDashboard(title, timezone)
	dbul.Overwrite = false
	body, err := json.Marshal(dbul)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	err = c.post(API_DASHBOARDS_DB, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return c.GetDashboard(title)
}

func (c *Client) GetDashboard(title string) (*Dashboard, error) {
	var db DashboardRes

	body, err := c.get(fmt.Sprintf("%s%s", API_DASHBOARDS_DB, SEOURL(title)))
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &db)
	if err != nil {
		fmt.Println(err)
	}
	body, _ = json.Marshal(db)
	fmt.Println(string(body))
	return db.Dashboard, err
}

func (c *Client) DeleteDashboard(title string) error {
	return c.delete(fmt.Sprintf("%s%s", API_DASHBOARDS_DB, SEOURL(title)))
}

func (c *Client) UploadDashboard(db *Dashboard) error {
	var dbul DashboardUploader
	dbul.Dashboard = db
	dbul.Overwrite = true

	body, err := json.Marshal(dbul)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return c.post(API_DASHBOARDS_DB, bytes.NewBuffer(body))
}

package main

import (
	"bytes"
)

type DashBoard struct {
	ID       int64          `json: "id"`
	Title    string         `json: "title"`
	Tags     []string       `json: "tags"`
	Timezone string         `json: "timezone"`
	Rows     []DashBoardRow `json: "rows"`
}

type DashBoardRow struct {
	Collapse     bool                 `json: "cololapse"`
	HideControls bool                 `json: "hideControls"`
	Height       string               `json: "height"`
	Pannels      []DashBoardRowPannel `json: "pannels"`
}

type DashBoardRowPannel struct {
}

func (c *Client) NewDashBoard(dbjson string) error {
	return c.post(API_DASHBOARDS_DB, bytes.NewBuffer([]byte(dbjson)))
}

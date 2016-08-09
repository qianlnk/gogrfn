package gografana

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (c *Client) GetDashBoard(slug string) (*DashBoard, error) {
	db := new(DashBoard)

	body, err := c.get(fmt.Sprintf("%s%s", API_DASHBOARDS_DB, slug))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, db)
	return db, err
}

func (c *Client) DeleteDashBoard(slug string) error {
	return c.delete(fmt.Sprintf("%s%s", API_DASHBOARDS_DB, slug))
}

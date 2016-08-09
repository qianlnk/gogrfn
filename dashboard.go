package gografana

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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

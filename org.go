package main

import (
	"bytes"
	"encoding/json"
	//"fmt"
)

type Org struct {
	ID   int64  `json: "id"`
	Name string `json: "name"`
	//Addr Address `json: address`
}

type Address struct {
	Address1 string `json: address1`
	Address2 string `json: address2`
	City     string `json: city`
	ZipCode  string `json: zipCode`
	State    string `json: state`
	Country  string `json: country`
}

func (c *Client) Orgs() ([]Org, error) {
	orgs := make([]Org, 0)

	body, err := c.get(API_ORGS)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &orgs)
	return orgs, err
}

func (c *Client) NewOrg(name string) error {
	orgs := map[string]string{
		"name": name,
	}

	data, err := json.Marshal(orgs)
	if err != nil {
		return err
	}

	return c.post(API_ORGS, bytes.NewBuffer(data))
}

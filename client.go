package main

import (
	"fmt"
	"net/http"
)

type Client struct {
	ApiKey  string
	BaseUrl string
	Cli     *http.Client
}

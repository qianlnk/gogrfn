package gografana

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	ApiKey   string
	BaseURL  *url.URL
	User     string
	Password string
	*http.Client
}

func NewGrafanaClient(key, baseURL, user, password string) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	u.User = url.UserPassword(user, password)

	if key == "" {
		return nil, errors.New("Api authorization not set!")
	}

	return &Client{
		fmt.Sprintf("Bearer %s", key),
		u,
		user,
		password,
		&http.Client{},
	}, nil

}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	u := c.BaseURL
	u.Path = path
	fmt.Println(u.String())
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func (c *Client) get(path string) ([]byte, error) {
	req, err := c.newRequest(METHOD_GET, path, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	mybody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("%s %s", res.Status, string(mybody)))
	}

	return mybody, nil

}
func (c *Client) post(path string, body io.Reader) error {
	req, err := c.newRequest(METHOD_POST, path, body)
	if err != nil {
		return err
	}

	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	mybody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("%s %s", res.Status, string(mybody)))
	}

	return nil
}

func (c *Client) delete(path string) error {
	req, err := c.newRequest(METHOD_DELETE, path, nil)
	if err != nil {
		return err
	}

	res, err := c.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	mybody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("%s %s", res.Status, string(mybody)))
	}

	return nil
}

//
//
//
//
//
//

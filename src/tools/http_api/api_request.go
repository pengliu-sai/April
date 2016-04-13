package http_api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type deadlinedConn struct {
	Timeout time.Duration
	net.Conn
}

func (c *deadlinedConn) Read(b []byte) (n int, err error) {
	c.Conn.SetReadDeadline(time.Now().Add(c.Timeout))
	return c.Conn.Read(b)
}

func (c *deadlinedConn) Write(b []byte) (n int, err error) {
	c.Conn.SetWriteDeadline(time.Now().Add(c.Timeout))
	return c.Conn.Write(b)
}

func NewDeadlineTransport(timeout time.Duration) *http.Transport {
	transport := &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, timeout)
			if err != nil {
				return nil, err
			}
			return &deadlinedConn{timeout, c}, nil
		},
	}
	return transport
}

type Client struct {
	c *http.Client
}

func NewClient(tlsConfig *tls.Config) *Client {
	transport := NewDeadlineTransport(2 * time.Second)
	transport.TLSClientConfig = tlsConfig
	return &Client{
		c: &http.Client{
			Transport: transport,
		},
	}
}

func (c *Client) GETV1(endpoint string, v interface{}) error {
retry:
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/vnd.liupeng; version=1.0")

	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		if resp.StatusCode == 403 && !strings.HasPrefix(endpoint, "https") {
			endpoint, err = httpsEndpoint(endpoint, body)
			if err != nil {
				return err
			}
			goto retry
		}
		return fmt.Errorf("got response %s %q", resp.StatusCode, body)
	}
	err = json.Unmarshal(body, &v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) POSTV1(endpoint string) error {
retry:
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/vnd.liupeng; version=1.0")

	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == 403 && !strings.HasPrefix(endpoint, "https") {
			endpoint, err = httpsEndpoint(endpoint, body)
			if err != nil {
				return err
			}
			goto retry
		}
		return fmt.Errorf("got response %s %q", resp.Status, body)
	}

	return nil
}

func httpsEndpoint(endpoint string, body []byte) (string, error) {
	var forbiddenResp struct {
		HTTPSPort int `json:"https_port"`
	}
	err := json.Unmarshal(body, &forbiddenResp)
	if err != nil {
		return "", err
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}

	host, _, err := net.SplitHostPort(u.Host)
	if err != nil {
		return "", err
	}

	u.Scheme = "https"
	u.Host = net.JoinHostPort(host, strconv.Itoa(forbiddenResp.HTTPSPort))
	return u.String(), nil
}

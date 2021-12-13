package nalo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Config struct {
	ApiBaseURL string
	Username   string
	Password   string
	Type       string
	Delivery   string
	Source     string
}

type Nalo struct {
	ctx     context.Context
	baseURL string
	params  url.Values
}

func New(ctx context.Context, cfg *Config) *Nalo {
	params := url.Values{}
	params.Set("username", cfg.Username)
	params.Set("password", cfg.Password)
	params.Set("type", cfg.Type)
	params.Set("dlr", cfg.Delivery)
	params.Set("source", cfg.Source)
	return &Nalo{
		ctx:     ctx,
		baseURL: cfg.ApiBaseURL,
		params:  params,
	}
}

func (n *Nalo) buildURL(message, destination string) string {
	n.params.Set("destination", destination)
	n.params.Set("message", message)
	return n.baseURL + n.params.Encode()
}

func (n *Nalo) SendSMS(message, destination string) error {
	url := n.buildURL(message, destination)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("request to %q failed with status '%d'", url, resp.StatusCode)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if splits := bytes.Split(b, []byte("|")); !bytes.Equal(splits[0], []byte("1701")) {
		return fmt.Errorf("external service failed: %s", "to send sms")
	}
	return nil
}

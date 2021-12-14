package nalo

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Nalo struct {
	ctx     context.Context
	baseURL string
	params  url.Values
	client  *http.Client
}

func New(ctx context.Context, username, password string, client *http.Client) *Nalo {
	if client == nil {
		client = http.DefaultClient
	}
	params := url.Values{}
	params.Set("username", username)
	params.Set("password", password)
	return &Nalo{
		ctx:     ctx,
		baseURL: "https://api.nalosolutions.com",
		params:  params,
		client:  client,
	}
}

func (n *Nalo) WithDelivery(dlr Delivery) *Nalo {
	n.params.Set("dlr", dlr.String())
	return n
}

func (n *Nalo) WithType(stype MessageType) *Nalo {
	n.params.Set("type", stype.String())
	return n
}

func (n *Nalo) buildURL(path, message, destination, from string) string {
	n.params.Set("destination", destination)
	n.params.Set("message", message)
	n.params.Set("source", from)
	return fmt.Sprintf("%s/%s?%s", n.baseURL, path, n.params.Encode())
}

func (n *Nalo) SendSMS(message, to, from string) error {
	url := n.buildURL("bulksms", message, to, from)
	resp, err := n.client.Get(url)
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

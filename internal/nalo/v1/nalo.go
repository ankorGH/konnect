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

func (n *Nalo) SendSMS(message, to, from string) (*Response, error) {
	url := n.buildURL("bulksms", message, to, from)
	resp, err := n.client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("request to %q failed with status '%d'", url, resp.StatusCode)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	responseSplit := bytes.Split(b, []byte("|"))
	switch string(responseSplit[0]) {
	case "1701":
		return NewResponse(string(responseSplit[1]), string(responseSplit[2])), nil
	case "1702":
		return nil, ErrInvalidURL
	case "1703":
		return nil, ErrInvalidCredentials
	case "1704":
		return nil, ErrInvalidType
	case "1705":
		return nil, ErrInvalidMessage
	case "1706":
		return nil, ErrInvalidDestination
	case "1707":
		return nil, ErrInvalidSource
	case "1708":
		return nil, ErrInvalidDLR
	case "1709":
		return nil, ErrInvalidUserValidation
	case "1710":
		return nil, ErrInternal
	case "1025":
		return nil, ErrInsufficientCreditUser
	case "1026":
		return nil, ErrInsufficientCreditReseller
	default:
		return nil, ErrUnknown
	}
}

package nalo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Credentials struct {
	Username string
	Password string
}

type Nalo struct {
	ctx        context.Context
	baseURL    string
	baseParams url.Values
	client     *http.Client
}

func New(ctx context.Context, credentials Credentials, client *http.Client) *Nalo {
	if client == nil {
		client = http.DefaultClient
	}
	baseParams := url.Values{}
	baseParams.Set("username", credentials.Username)
	baseParams.Set("password", credentials.Password)
	return &Nalo{
		ctx:        ctx,
		baseURL:    "https://api.nalosolutions.com",
		baseParams: baseParams,
		client:     client,
	}
}

func (n *Nalo) buildSMSURL(message, destination, from string, delivery Delivery, messageType MessageType) string {
	params := url.Values{}
	params.Set("destination", destination)
	params.Set("message", message)
	params.Set("source", from)
	params.Set("dlr", delivery.String())
	params.Set("type", messageType.String())
	return fmt.Sprintf("%s/bulksms?%s&%s", n.baseURL, params.Encode(), n.baseParams.Encode())
}

func (n *Nalo) GetBalance() (*CreditBalanceResponse, error) {
	url := fmt.Sprintf("%s/nalosms/credit_bal.php?%s", n.baseURL, n.baseParams.Encode())
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
	var creditBalance CreditBalanceResponse
	if err := json.Unmarshal(b, &creditBalance); err != nil {
		return nil, err
	}
	return &creditBalance, nil
}

func (n *Nalo) SendSMS(message, to, from string, delivery Delivery, messageType MessageType) (*SMSResponse, error) {
	url := n.buildSMSURL(message, to, from, delivery, messageType)
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
		return &SMSResponse{
			Destination: string(responseSplit[1]),
			MessageId:   string(responseSplit[2]),
		}, nil
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

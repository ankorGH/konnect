package nalo

type Response struct {
	Destination string
	MessageId   string
}

func NewResponse(destination, messageId string) *Response {
	return &Response{
		Destination: destination,
		MessageId:   messageId,
	}
}

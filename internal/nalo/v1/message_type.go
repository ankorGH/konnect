package nalo

import (
	"fmt"
	"strconv"
	"strings"
)

type MessageType int64

const (
	MessageType_PlainTextGSM MessageType = 0 // (GSM 3.38 Character encoding)
	MessageType_FlashGSM     MessageType = 1 // (GSM 3.38 Character encoding)
	MessageType_Unicode      MessageType = 2
	MessageType_Reserved     MessageType = 3
	MessageType_WAPPush      MessageType = 4
	MessageType_PlainTextISO MessageType = 5 // (ISO-8859-1 Character encoding)
	MessageType_UnicodeFlash MessageType = 6
	MessageType_FlashISO     MessageType = 7 // (ISO-8859-1 Character encoding)
	Message_Invalid          MessageType = 1001
)

func ParseMessageType(s string) MessageType {
	switch strings.TrimSpace(s) {
	case "0":
		return MessageType_PlainTextGSM
	case "1":
		return MessageType_FlashGSM
	case "2":
		return MessageType_Unicode
	case "3":
		return MessageType_Reserved
	case "4":
		return MessageType_WAPPush
	case "5":
		return MessageType_UnicodeFlash
	case "6":
		return MessageType_UnicodeFlash
	case "7":
		return MessageType_FlashISO
	default:
		return Message_Invalid
	}
}

func (m MessageType) String() string {
	return strconv.FormatInt(int64(m), 10)
}

func (m MessageType) IsValid() bool {
	return m != Message_Invalid
}

func (m MessageType) MarshalText() (int64, error) {
	return int64(m), nil
}

func (m *MessageType) UnmarshalText(data int64) error {
	v := ParseMessageType(strconv.FormatInt(data, 10))
	if v.IsValid() {
		return fmt.Errorf("invalid message type '%d'", data)
	}
	*m = v
	return nil
}

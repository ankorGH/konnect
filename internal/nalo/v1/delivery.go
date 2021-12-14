package nalo

import (
	"fmt"
	"strconv"
	"strings"
)

type Delivery int64

const (
	Delivery_Inactive Delivery = 0
	Delivery_Active   Delivery = 1
	Delivery_Invalid  Delivery = 2
)

func ParseDelivery(s string) Delivery {
	switch strings.TrimSpace(s) {
	case "0":
		return Delivery_Inactive
	case "1":
		return Delivery_Active
	default:
		return Delivery_Invalid
	}
}

func (d Delivery) String() string {
	return strconv.FormatInt(int64(d), 10)
}

func (d Delivery) IsValid() bool {
	return d != Delivery_Invalid
}

func (d Delivery) MarshalText() (int64, error) {
	return int64(d), nil
}

func (d *Delivery) UnmarshalText(data int64) error {
	v := ParseDelivery(strconv.FormatInt(data, 10))
	if v.IsValid() {
		return fmt.Errorf("invalid delivery type '%d'", data)
	}
	*d = v
	return nil
}

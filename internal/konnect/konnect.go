package konnect

import (
	"context"
	"fmt"
	"time"

	"github.com/ankorgh/konnect/internal/nalo/v1"
)

type Config struct {
	Interval  time.Duration
	Messages  []string
	Day       time.Time
	StartHour int
	Sender    string
	Source    string
}

type Konnect struct {
	cfg      *Config
	ctx      context.Context
	notifier *nalo.Nalo
}

func New(ctx context.Context, cfg *Config, notifier *nalo.Nalo) *Konnect {
	return &Konnect{
		ctx:      ctx,
		cfg:      cfg,
		notifier: notifier,
	}
}

func (k *Konnect) Run() error {
	msgsLen := len(k.cfg.Messages)
	if msgsLen < 1 {
		return fmt.Errorf("%v", "no messages available")
	}
	if time.Now().UTC().After(k.cfg.Day) {
		return fmt.Errorf("messages were supposed to be sent before: %q", k.cfg.Day)
	}
	for {
		if time.Now().UTC().Hour() < k.cfg.StartHour {
			time.Sleep(time.Minute * 10)
		} else {
			break
		}
	}
	msgIdx := 0
	ticker := time.NewTicker(k.cfg.Interval)
	for {
		if time.Now().UTC().After(k.cfg.Day) {
			return fmt.Errorf("messages were supposed to be sent before: %q", k.cfg.Day)
		}
		if msgsLen <= msgIdx {
			break
		}
		select {
		case <-ticker.C:
			_, err := k.notifier.SendSMS(k.cfg.Messages[msgIdx], k.cfg.Sender, k.cfg.Source, nalo.Delivery_Active, nalo.MessageType_PlainTextISO)
			if err != nil {
				return err
			}
			msgIdx++
		case <-k.ctx.Done():
			return fmt.Errorf("context shutdown")
		}
	}
	return nil
}

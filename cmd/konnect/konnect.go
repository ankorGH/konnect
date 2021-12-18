package main

import (
	"context"
	"fmt"

	"github.com/ankorgh/konnect/internal/config"
	"github.com/ankorgh/konnect/internal/konnect"
	"github.com/ankorgh/konnect/internal/nalo/v1"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	ctx := context.Background()
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	// setup nalo sms api
	notifier := nalo.New(ctx, nalo.Credentials{
		Username: cfg.GetString("username"),
		Password: cfg.GetString("password"),
	}, nil)

	// setup konnect
	kCfg := &konnect.Config{
		Interval:  cfg.GetDuration("interval"),
		Messages:  cfg.GetStringSlice("messages"),
		Day:       cfg.GetTime("day").UTC(),
		StartHour: cfg.GetInt("startHour"),
		Sender:    cfg.GetString("receiver"),
		Source:    cfg.GetString("source"),
	}
	konnect := konnect.New(ctx, kCfg, notifier)
	err = konnect.Run()
	if err != nil {
		return err
	}

	return nil
}

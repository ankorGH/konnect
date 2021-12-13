package main

import (
	"context"
	"fmt"

	"github.com/ankorgh/konnect/internal/config"
	"github.com/ankorgh/konnect/internal/konnect"
	"github.com/ankorgh/konnect/internal/nalo"
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
	naloCfg := nalo.Config{
		ApiBaseURL: cfg.GetString("apiBaseURL"),
		Username:   cfg.GetString("username"),
		Password:   cfg.GetString("password"),
		Type:       cfg.GetString("type"),
		Delivery:   cfg.GetString("delivery"),
		Source:     cfg.GetString("source"),
	}
	notifier := nalo.New(ctx, &naloCfg)

	// setup konnect
	kCfg := &konnect.Config{
		Interval:  cfg.GetDuration("interval"),
		Messages:  cfg.GetStringSlice("messages"),
		Day:       cfg.GetTime("day").UTC(),
		StartHour: cfg.GetInt("startHour"),
		Sender:    cfg.GetString("receiver"),
	}
	konnect := konnect.New(ctx, kCfg, notifier)
	err = konnect.Run()
	if err != nil {
		return err
	}

	return nil
}

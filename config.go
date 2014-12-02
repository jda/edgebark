package main

import (
	"errors"
	"log"
	"os"
)

type config struct {
	Port        string // yes, is properly a int but we always use it as a str so...
	HipchatKey  string
	HipchatRoom string
}

func loadConfig() error {
	inError := false

	cfg.Port = os.Getenv("PORT")
	if cfg.Port == "" {
		inError = true
		log.Println("config error: PORT not set")
	}

	cfg.HipchatKey = os.Getenv("HIPCHAT_KEY")
	if cfg.HipchatKey == "" {
		inError = true
		log.Println("config error: HIPCHAT_KEY not set")
	}

	cfg.HipchatRoom = os.Getenv("HIPCHAT_ROOM")
	if cfg.HipchatRoom == "" {
		inError = true
		log.Println("config error: HIPCHAT_ROOM not set")
	}

	if inError {
		return errors.New("config env incomplete")
	}

	return nil
}
